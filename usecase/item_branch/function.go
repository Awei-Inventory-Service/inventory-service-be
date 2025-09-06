package item_branch

import (
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *itemBranchUsecase) FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchDomain.FindByBranchAndItem(payload.BranchId, payload.ItemId)
}

func (s *itemBranchUsecase) FindByBranchId(branchId string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchDomain.FindByBranch(branchId)
}

func (s *itemBranchUsecase) FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	return s.itemBranchDomain.FindAll()
}
