package item_branch

import (
	itemBranch "github.com/inventory-service/domain/item_branch"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ItemBranchUsecase interface {
	FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByBranchId(branchId string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
}

type itemBranchUsecase struct {
	itemBranchDomain itemBranch.ItemBranchDomain
}
