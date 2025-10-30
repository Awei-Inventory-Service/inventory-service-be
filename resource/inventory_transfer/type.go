package inventory_transfer

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type InventoryTransferResource interface {
	Create(ctx context.Context, payload model.InventoryTransfer) (model.InventoryTransfer, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, id string, payload model.InventoryTransfer) (model.InventoryTransfer, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.InventoryTransfer, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, id string) (model.InventoryTransfer, *error_wrapper.ErrorWrapper)
	UpdateStatus(ctx context.Context, id, status string) (errW *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, payload model.InventoryTransfer) (errW *error_wrapper.ErrorWrapper)
}

type inventoryTransferResource struct {
	db *gorm.DB
}
