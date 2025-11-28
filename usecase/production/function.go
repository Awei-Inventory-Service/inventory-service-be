package production

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/constant"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
)

func (p *productionUsecase) Create(ctx context.Context, payload dto.CreateProductionRequest) (*model.Production, *error_wrapper.ErrorWrapper) {
	var (
		updatedBranchItems []string
		totalCostMovement  = 0.0
	)

	productionParsedDate, err := time.Parse("2006-01-02", payload.ProductionDate)
	if err != nil {
		return nil, error_wrapper.New(model.ErrInvalidTimestamp, err.Error())
	}

	for _, sourceItem := range payload.SourceItems {
		initialInventory, errW := p.inventoryDomain.FindByBranchAndItem(payload.BranchID, sourceItem.SourceItemID)

		if errW != nil {
			fmt.Println("INi error", errW)
			return nil, errW
		}
		payloadQuantity := utils.StandarizeMeasurement(sourceItem.InitialQuantity, sourceItem.InitialUnit, initialInventory.Item.Unit)
		if initialInventory.Stock < payloadQuantity {
			return nil, error_wrapper.New(model.UErrStockIsNotEnough, fmt.Sprintf("Current inventory stock: %f is not enough for : %f", initialInventory.Stock, sourceItem.InitialQuantity))
		}
	}

	production, errW := p.productionDomain.Create(ctx, payload)

	if errW != nil {
		return nil, errW
	}

	// Create stock transcation in for the final
	referenceType := constant.Production

	for _, productionItem := range payload.SourceItems {
		item, errW := p.itemDomain.FindByID(ctx, productionItem.SourceItemID)
		if errW != nil {
			fmt.Println("Error finding item by id", errW)
			continue
		}

		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, dto.CustomDate{
			Day:   productionParsedDate.Day(),
			Month: int(productionParsedDate.Month()),
			Year:  productionParsedDate.Year(),
		}, productionItem.SourceItemID, payload.BranchID)
		if errW != nil {
			fmt.Println("Error getting inventory by date", errW)
			continue
		}

		standarizedQuantity := utils.StandarizeMeasurement(productionItem.InitialQuantity, productionItem.InitialUnit, item.Unit)
		cost := inventory.Price * standarizedQuantity

		errW = p.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              productionItem.SourceItemID,
			Type:                "OUT",
			Quantity:            productionItem.InitialQuantity,
			Unit:                productionItem.InitialUnit,
			IssuerID:            payload.UserID,
			Reference:           production.UUID,
			ReferenceType:       &referenceType,
			Cost:                cost,
			TransactionDate:     productionParsedDate,
		})

		totalCostMovement += cost
		updatedBranchItems = append(updatedBranchItems, productionItem.SourceItemID)
		if errW != nil {
			return nil, errW
		}

		errW = p.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			BranchID: payload.BranchID,
			ItemID:   productionItem.SourceItemID,
			NewTime:  payload.ProductionDate,
		})
		if errW != nil {
			fmt.Println("Error recalculating inventory", errW)
			return nil, errW
		}
	}

	errW = p.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.FinalItemID,
		Type:                "IN",
		Quantity:            payload.FinalQuantity,
		IssuerID:            payload.UserID,
		Unit:                payload.FinalUnit,
		Reference:           production.UUID,
		ReferenceType:       &referenceType,
		Cost:                totalCostMovement,
		TransactionDate:     productionParsedDate,
	})

	if errW != nil {
		return nil, errW
	}
	updatedBranchItems = append(updatedBranchItems, payload.FinalItemID)
	errW = p.inventoryDomain.BulkSyncBranchItems(ctx, payload.BranchID, updatedBranchItems)

	if errW != nil {
		return nil, errW
	}

	return production, nil
}

func (p *productionUsecase) Get(ctx context.Context, filter dto.GetProductionFilter) ([]dto.GetProductionList, *error_wrapper.ErrorWrapper) {
	return p.productionDomain.Get(ctx, filter)
}

func (p *productionUsecase) Delete(ctx context.Context, payload dto.DeleteProductionRequest) *error_wrapper.ErrorWrapper {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	production, errW := p.productionDomain.Get(ctx, dto.GetProductionFilter{
		ProductionID: payload.ProductionID,
		BranchID:     payload.BranchID,
	})
	if errW != nil {
		fmt.Println("Error getting production ", errW)
		return errW
	}

	if len(production) == 0 {
		return error_wrapper.New(model.RErrDataNotFound, "Production not found")
	}

	deletedProduction := production[0]

	errW = p.productionDomain.Delete(ctx, payload.ProductionID)
	if errW != nil {
		return errW
	}

	_, errW = p.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": payload.ProductionID,
		},
	}, payload.UserID)
	if errW != nil {
		fmt.Println("Error invalidating stock transaction", errW)
		return errW
	}

	for _, item := range deletedProduction.SourceItems {
		// Recalculate inventory for specific item
		errW = p.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			BranchID: payload.BranchID,
			ItemID:   item.SourceItemID,
			NewTime:  deletedProduction.ProductionDate,
		})
		if errW != nil {
			fmt.Println("Item: ", item)
			fmt.Println("Error recalculating inventory", errW)
			continue
		}
		_, _, errW = p.inventoryDomain.SyncBranchItem(ctx, deletedProduction.BranchID, item.SourceItemID)
		if errW != nil {
			fmt.Println("Error syncing branch item", errW)
			continue
		}
	}

	return nil
}

func (p *productionUsecase) Update(ctx context.Context, payload dto.UpdateProductionRequest) (model.Production, *error_wrapper.ErrorWrapper) {

	var (
		affectedItems []string
	)
	// 1. Validate payload
	productionDate, err := time.Parse("2006-01-02", payload.ProductionDate)
	if err != nil {
		return model.Production{}, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
	}
	valid, errW := p.ValidateSourceItemsQuantity(ctx, payload.SourceItems, productionDate, payload.BranchID)
	if errW != nil || !valid {
		return model.Production{}, errW
	}

	// Update the production data
	errW = p.productionDomain.Update(ctx, payload)
	if errW != nil {
		fmt.Println("Error updating production domain", errW)
		return model.Production{}, errW
	}

	// 1. Delete all production_item
	errW = p.productionItemDomain.Delete(ctx, model.ProductionItem{
		ProductionID: payload.ProductionID,
	})
	if errW != nil {
		fmt.Println("Error deleting production", errW)
		return model.Production{}, errW
	}

	// 2. Delete all stock_transactions for production
	oldItems, errW := p.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": payload.ProductionID,
		},
	}, payload.UserID)
	if errW != nil {
		fmt.Println("Error invalidating stock transactions", errW)
		return model.Production{}, errW
	}
	affectedItems = append(affectedItems, oldItems...)

	// 3. Create new stock_transaction
	totalCost := 0.0
	for _, productionItem := range payload.SourceItems {
		item, errW := p.itemDomain.FindByID(ctx, productionItem.SourceItemID)
		if errW != nil {
			fmt.Println("Error finding item by id", errW)
			continue
		}

		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, dto.CustomDate{
			Day:   productionDate.Day(),
			Month: int(productionDate.Month()),
			Year:  productionDate.Year(),
		}, productionItem.SourceItemID, payload.BranchID)
		if errW != nil {
			fmt.Println("Error getting inventory by date", errW)
			return model.Production{}, errW
		}

		standarizedQuantity := utils.StandarizeMeasurement(productionItem.InitialQuantity, productionItem.InitialUnit, item.Unit)
		cost := inventory.Price * standarizedQuantity
		// Create stock transaction
		errW = p.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              item.UUID,
			Type:                "OUT",
			Quantity:            productionItem.InitialQuantity,
			Unit:                productionItem.InitialUnit,
			IssuerID:            payload.UserID,
			Reference:           payload.ProductionID,
			Cost:                cost,
			TransactionDate:     productionDate,
		})
		if errW != nil {
			fmt.Println("Error creating new stock transaction", errW)
			continue
		}

		waste := standarizedQuantity - payload.FinalQuantity

		// Create new production item
		_, errW = p.productionItemDomain.Create(ctx, model.ProductionItem{
			ProductionID:    payload.ProductionID,
			SourceItemID:    productionItem.SourceItemID,
			Quantity:        productionItem.InitialQuantity,
			Unit:            productionItem.InitialUnit,
			WasteQuantity:   waste,
			WastePercentage: waste / productionItem.InitialQuantity * 100,
		})
		if errW != nil {
			fmt.Println("Error creating new production item ", errW)
			continue
		}

		totalCost += cost
		affectedItems = append(affectedItems, productionItem.SourceItemID)
	}

	errW = p.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.FinalItemID,
		Type:                "IN",
		Quantity:            payload.FinalQuantity,
		IssuerID:            payload.UserID,
		Unit:                payload.FinalUnit,
		Reference:           payload.ProductionID,
		Cost:                totalCost,
		TransactionDate:     productionDate,
	})
	if errW != nil {
		fmt.Println("Error creating stock transaction for in type", errW)
		return model.Production{}, errW
	}
	affectedItems = append(affectedItems, payload.FinalItemID)
	// Sync inventory for all affected items
	for _, item := range affectedItems {
		errW = p.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			BranchID: payload.BranchID,
			ItemID:   item,
			NewTime:  payload.ProductionDate,
		})
		if errW != nil {
			fmt.Println("Error recalculating inventory", errW)
			continue
		}
	}
	errW = p.inventoryDomain.BulkSyncBranchItems(ctx, payload.BranchID, affectedItems)
	if errW != nil {
		fmt.Println("Error doing bulk sync branch item", errW)
		return model.Production{}, errW
	}

	return model.Production{}, nil
}

func (p *productionUsecase) ValidateSourceItemsQuantity(ctx context.Context, sourceItems []dto.SourceItemCreateProductionRequest, productionDate time.Time, branchID string) (valid bool, errW *error_wrapper.ErrorWrapper) {

	for _, newItem := range sourceItems {
		item, errW := p.itemDomain.FindByID(ctx, newItem.SourceItemID)
		if errW != nil {
			fmt.Println("Error getting item with id", newItem.SourceItemID)
			fmt.Println("Error getting item by id", errW)
			continue
		}

		//A. Get item price at the production date
		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, dto.CustomDate{
			Day:   productionDate.Day(),
			Month: int(productionDate.Month()),
			Year:  productionDate.Year(),
		}, newItem.SourceItemID, branchID)

		if errW != nil {
			fmt.Println("Error getting item price", errW)
			continue
		}

		// B. Standarize quantity
		payloadQuantity := utils.StandarizeMeasurement(newItem.InitialQuantity, newItem.InitialUnit, item.Unit)
		if inventory.Balance < payloadQuantity {
			errW = error_wrapper.New(model.UErrStockIsNotEnough, fmt.Sprintf("Item : %s balance is only %f on %s", item.Name, inventory.Balance, productionDate))
			return false, errW
		}
	}

	return true, nil
}
