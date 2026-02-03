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

		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, productionParsedDate, productionItem.SourceItemID, payload.BranchID)
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

	errW = p.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
		ItemID:   payload.FinalItemID,
		BranchID: payload.BranchID,
		NewTime:  payload.ProductionDate,
	})
	if errW != nil {
		fmt.Println("Error recalculating inventory", err)
		return nil, errW
	}

	return production, nil
}

func (p *productionUsecase) Get(ctx context.Context, payload dto.GetListRequest) (dto.GetProductionResponse, *error_wrapper.ErrorWrapper) {
	return p.productionDomain.Get(ctx, payload)
}

func (p *productionUsecase) Delete(ctx context.Context, payload dto.DeleteProductionRequest) *error_wrapper.ErrorWrapper {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	filter := []dto.Filter{{Key: "uuid", Values: []string{payload.ProductionID}}}
	production, errW := p.productionDomain.Get(ctx, dto.GetListRequest{
		Filter: filter,
		Limit:  0,
		Offset: 0,
	})
	if errW != nil {
		fmt.Println("Error getting production ", errW)
		return errW
	}

	if len(production.Productions) == 0 {
		return error_wrapper.New(model.RErrDataNotFound, "Production not found")
	}

	deletedProduction := production.Productions[0]

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
	payload.ParsedProductionDate = productionDate

	oldProduction, errW := p.productionDomain.Get(ctx, dto.GetListRequest{
		Filter: []dto.Filter{{Key: "uuid", Values: []string{payload.ProductionID}}},
		Limit:  1,
	})
	if errW != nil {
		fmt.Println("Error getting old production ", errW)
		return model.Production{}, errW
	}

	if len(oldProduction.Productions) == 0 {
		errW = error_wrapper.New(model.RErrDataNotFound, fmt.Sprintf("Production with id %s not found ", payload.ProductionID))
		return model.Production{}, errW
	}

	oldProductionData := oldProduction.Productions[0]
	oldProductionDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", oldProductionData.ProductionDate)
	if err != nil {
		fmt.Println("Invalid time stamp for old production date", oldProductionData.ProductionDate)
		return model.Production{}, error_wrapper.New(model.ErrInvalidTimestamp, "invalid start time format: "+err.Error())
	}

	valid, errW := p.ValidateSourceItemsQuantity(ctx, payload, oldProductionData)
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
	productionType := constant.Production
	for _, productionItem := range payload.SourceItems {
		item, errW := p.itemDomain.FindByID(ctx, productionItem.SourceItemID)
		if errW != nil {
			fmt.Println("Error finding item by id", errW)
			continue
		}

		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, productionDate, productionItem.SourceItemID, payload.BranchID)
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
			ReferenceType:       &productionType,
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
		ReferenceType:       &productionType,
	})
	if errW != nil {
		fmt.Println("Error creating stock transaction for in type", errW)
		return model.Production{}, errW
	}
	affectedItems = append(affectedItems, payload.FinalItemID)
	// Sync inventory for all affected items
	for _, item := range affectedItems {
		errW = p.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			BranchID:     payload.BranchID,
			ItemID:       item,
			NewTime:      payload.ProductionDate,
			PreviousTime: &oldProductionDate,
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

func (p *productionUsecase) ValidateSourceItemsQuantity(ctx context.Context, payload dto.UpdateProductionRequest, oldProduction dto.GetProduction) (valid bool, errW *error_wrapper.ErrorWrapper) {
	var (
		mappedProductionSourceItems = make(map[string]dto.GetProductionItem)
	)

	for _, oldItem := range oldProduction.SourceItems {
		mappedProductionSourceItems[oldItem.SourceItemID] = oldItem
	}

	for _, newItem := range payload.SourceItems {
		var (
			stockInProduction = 0.0
			unit              = ""
		)

		item, errW := p.itemDomain.FindByID(ctx, newItem.SourceItemID)
		if errW != nil {
			fmt.Println("Error getting item with id", newItem.SourceItemID)
			fmt.Println("Error getting item by id", errW)
			continue
		}

		unit = item.Unit
		inventory, errW := p.inventoryDomain.GetInventoryByDate(ctx, payload.ParsedProductionDate, newItem.SourceItemID, payload.BranchID)
		if errW != nil {
			fmt.Println("Error getting item price", errW)
			continue
		}

		if oldValue, existInOldProduction := mappedProductionSourceItems[newItem.SourceItemID]; existInOldProduction {
			// maximumStock += oldValue.Quantity
			stockInProduction += utils.StandarizeMeasurement(oldValue.InitialQuantity, oldValue.SourceItemUnit, unit)
		}
		maximumStock := stockInProduction + inventory.Balance
		payloadQuantity := utils.StandarizeMeasurement(newItem.InitialQuantity, newItem.InitialUnit, item.Unit)

		if maximumStock < payloadQuantity {
			errW = error_wrapper.New(model.UErrStockIsNotEnough, fmt.Sprintf("Item: %s balance is only %f on %s", item.Name, maximumStock, payload.ProductionDate))
			return false, errW
		}
	}

	return true, nil
}

func (p *productionUsecase) GetByID(ctx context.Context, id string) (production dto.GetProduction, errW *error_wrapper.ErrorWrapper) {
	filter := []dto.Filter{
		{
			Key:    "uuid",
			Values: []string{id},
		},
	}

	productionRaw, errW := p.productionDomain.Get(ctx, dto.GetListRequest{
		Filter: filter,
		Limit:  0,
		Offset: 0,
	})

	production = productionRaw.Productions[0]
	return
}
