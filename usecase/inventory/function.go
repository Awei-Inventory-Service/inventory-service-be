package inventory

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryUsecase) Create(ctx context.Context, payload dto.CreateInventoryRequest) *error_wrapper.ErrorWrapper {

	// 1. Create out transaction for the item compositions
	item, errW := i.itemDomain.FindByID(ctx, payload.ItemID)

	if errW != nil {
		return errW
	}
	referenceType := "ITEM_CREATION"

	for _, itemComposition := range item.ChildCompositions {
		fmt.Println("iNI ITEM COMPOSIITON ITEM ID", itemComposition.ChildItemID)
		// total := itemComposition.Ratio * payload.Quantity * itemComposition.PortionSize
		errW := i.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              itemComposition.ChildItemID,
			Type:                "OUT",
			IssuerID:            payload.UserID,
			Quantity:            0,
			Cost:                0.0,
			Unit:                itemComposition.Unit,
			Reference:           "",
			ReferenceType:       &referenceType,
		})

		if errW != nil {
			return errW
		}

		_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, itemComposition.ChildItemID)

		if errW != nil {
			return errW
		}

	}
	// 2. Create the inside transactions for the item

	errW = i.stockTransactionDomain.Create(model.StockTransaction{
		BranchOriginID:      payload.BranchID,
		BranchDestinationID: payload.BranchID,
		ItemID:              payload.ItemID,
		Type:                "IN",
		IssuerID:            payload.UserID,
		Quantity:            payload.Quantity,
		Unit:                item.Unit,
		Reference:           "",
		ReferenceType:       &referenceType,
	})

	if errW != nil {
		return errW
	}

	_, _, errW = i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)

	if errW != nil {
		return errW
	}

	return nil
}

func (i *inventoryUsecase) FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindByBranchAndItem(payload.BranchId, payload.ItemId)
}

func (i *inventoryUsecase) FindByBranchId(branchId string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindByBranch(branchId)
}

func (i *inventoryUsecase) FindAll() ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.FindAll()
}

func (i *inventoryUsecase) SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (currentStock, currentPrice float64, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)
}

func (i *inventoryUsecase) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.Get(ctx, filter, order, limit, offset)
}

func (i *inventoryUsecase) RecalculateInventory(ctx context.Context, payload dto.RecalculateInventoryRequest) (errW *error_wrapper.ErrorWrapper) {
	return i.inventoryDomain.RecalculateInventory(ctx, payload)
}
