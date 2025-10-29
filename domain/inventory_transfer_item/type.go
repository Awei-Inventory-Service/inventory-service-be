package inventory_transfer_item

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/inventory_transfer_item"
)

type InventoryTransferItemDomain interface {
	Create(ctx context.Context, payload model.InventoryTransferItem) (result model.InventoryTransferItem, errW *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, payload model.InventoryTransferItem) (errW *error_wrapper.ErrorWrapper)
}

type inventoryTransferItemDomain struct {
	inventoryTransferItemResource inventory_transfer_item.InventoryTransferItemResource
}
