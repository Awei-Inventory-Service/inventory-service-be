package branch_item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type BranchItemResource interface {
	Create(stockBalance model.BranchItem) *error_wrapper.ErrorWrapper
	FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
}

type branchItemResource struct {
	db *gorm.DB
}
