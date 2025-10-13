package production

import (
	"context"
	"fmt"

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
		inventory, errW := p.inventoryDomain.FindByBranchAndItem(payload.BranchID, productionItem.SourceItemID)

		if errW != nil {
			return nil, errW
		}
		cost := inventory.Value * productionItem.InitialQuantity

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
		})

		totalCostMovement += cost
		updatedBranchItems = append(updatedBranchItems, productionItem.SourceItemID)
		if errW != nil {
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

	errW = p.productionDomain.Delete(ctx, payload.ProductionID)

	if errW != nil {
		return errW
	}

	stockTransactions, errW := p.stockTransactionDomain.FindWithFilter([]map[string]interface{}{
		{
			"field": "reference",
			"value": payload.ProductionID,
		},
	}, "", 0, 0)

	if errW != nil {
		return errW
	}
	fmt.Println("INi stock transactions", stockTransactions)
	return nil
}
