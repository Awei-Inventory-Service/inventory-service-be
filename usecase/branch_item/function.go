package branch_item

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *branchItemUsecase) Create(ctx context.Context, payload dto.CreateBranchItemRequest) *error_wrapper.ErrorWrapper {

	// 1. Create out transaction for the item compositions
	item, errW := s.itemDomain.FindByID(ctx, payload.ItemID)

	if errW != nil {
		return errW
	}
	referenceType := "ITEM_CREATION"

	for _, itemComposition := range item.ChildCompositions {
		fmt.Println("iNI ITEM COMPOSIITON ITEM ID", itemComposition.ChildItemID)
		// total := itemComposition.Ratio * payload.Quantity * itemComposition.PortionSize
		errW := s.stockTransactionDomain.Create(model.StockTransaction{
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

		errW = s.branchItemDomain.SyncBranchItem(ctx, payload.BranchID, itemComposition.ChildItemID)

		if errW != nil {
			return errW
		}

	}
	// 2. Create the inside transactions for the item

	errW = s.stockTransactionDomain.Create(model.StockTransaction{
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

	errW = s.branchItemDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)

	if errW != nil {
		return errW
	}

	return nil
}

func (s *branchItemUsecase) FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemDomain.FindByBranchAndItem(payload.BranchId, payload.ItemId)
}

func (s *branchItemUsecase) FindByBranchId(branchId string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemDomain.FindByBranch(branchId)
}

func (s *branchItemUsecase) FindAll() ([]dto.GetBranchItemResponse, *error_wrapper.ErrorWrapper) {
	return s.branchItemDomain.FindAll()
}

func (b *branchItemUsecase) SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (errW *error_wrapper.ErrorWrapper) {
	return b.branchItemDomain.SyncBranchItem(ctx, payload.BranchID, payload.ItemID)
}
