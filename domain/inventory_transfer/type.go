package inventory_transfer

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/inventory_transfer"
	"github.com/inventory-service/resource/inventory_transfer_item"
)

type InventoryTransferDomain interface {
	Create(ctx context.Context, payload dto.CreateInventoryTransferRequest) (model.InventoryTransfer, *error_wrapper.ErrorWrapper)
	UpdateStatus(ctx context.Context, id, status string) (errW *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, id string) (result model.InventoryTransfer, errW *error_wrapper.ErrorWrapper)
}

type inventoryTransferDomain struct {
	inventoryTransferResource     inventory_transfer.InventoryTransferResource
	inventoryTransferItemResource inventory_transfer_item.InventoryTransferItemResource
}
