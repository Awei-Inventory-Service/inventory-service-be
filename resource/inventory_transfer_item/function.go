package inventory_transfer_item

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryTransferItemResource) Create(ctx context.Context, payload model.InventoryTransferItem) (newData model.InventoryTransferItem, errW *error_wrapper.ErrorWrapper) {
	result := i.db.WithContext(ctx).Create(&payload)

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
		return
	}

	return payload, errW
}

func (i *inventoryTransferItemResource) Delete(ctx context.Context, payload model.InventoryTransferItem) (errW *error_wrapper.ErrorWrapper) {
	query := i.db.WithContext(ctx)

	if payload.ItemID != "" {
		query = query.Where("item_id = ?", payload.ItemID)
	}

	if payload.InventoryTransferID != "" {
		query = query.Where("inventory_transfer_id = ?", payload.InventoryTransferID)
	}

	if payload.UUID != "" {
		query = query.Where("uuid = ? ", payload.UUID)
	}

	result := query.Delete(&model.InventoryTransferItem{})

	if result.Error != nil {
		errW = error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
		return
	}

	return
}
