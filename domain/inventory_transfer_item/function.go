package inventory_transfer_item

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryTransferItemDomain) Create(ctx context.Context, payload model.InventoryTransferItem) (result model.InventoryTransferItem, errW *error_wrapper.ErrorWrapper) {
	return i.inventoryTransferItemResource.Create(ctx, payload)
}

func (i *inventoryTransferItemDomain) Delete(ctx context.Context, payload model.InventoryTransferItem) (errW *error_wrapper.ErrorWrapper){
	return i.inventoryTransferItemResource.Delete(ctx, payload)
}
