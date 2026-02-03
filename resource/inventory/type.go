package inventory

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type InventoryResource interface {
	Create(inventory model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.Inventory, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper)
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (inventories []model.Inventory, count int64, errW *error_wrapper.ErrorWrapper)
}

type inventoryResource struct {
	db *gorm.DB
}
