package item_branch

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ItemBranchResource interface {
	Create(stockBalance model.ItemBranch) *error_wrapper.ErrorWrapper
	FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.ItemBranch, *error_wrapper.ErrorWrapper)
	Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
}

type itemBranchResource struct {
	db *gorm.DB
}
