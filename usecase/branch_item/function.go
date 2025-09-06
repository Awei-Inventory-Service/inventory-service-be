package branch_item

import (
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

func (s *branchItemUsecase) FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	return s.branchItemDomain.FindAll()
}
