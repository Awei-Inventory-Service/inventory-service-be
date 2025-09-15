package branch_item

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

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
	currentBalance, errW := b.branchItemDomain.SyncCurrentBalance(ctx, payload.BranchID, payload.ItemID)

	if errW != nil {
		return
	}

	currentPrice, errW := b.branchItemDomain.CalculatePrice(ctx, payload.BranchID, payload.ItemID, currentBalance)

	fmt.Println("INi current price", currentPrice)
	if errW != nil {
		return
	}

	newBranchItem := model.BranchItem{
		BranchID:     payload.BranchID,
		ItemID:       payload.ItemID,
		CurrentStock: currentBalance,
		Price:        currentPrice,
	}

	_, errW = b.branchItemDomain.Update(ctx, newBranchItem)

	if errW != nil {
		return
	}

	return
}
