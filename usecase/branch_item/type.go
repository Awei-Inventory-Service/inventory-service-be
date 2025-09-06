package branch_item

import (
	branchitem "github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type BranchItemUsecase interface {
	FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranchId(branchId string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper)
}

type branchItemUsecase struct {
	branchItemDomain branchitem.BranchItemDomain
}
