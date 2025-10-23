package inventory_transfer_item

import "gorm.io/gorm"

func NewInventoryTransferItemResource(
	db *gorm.DB,
) InventoryTransferItemResource {
	return &inventoryTransferItemResource{
		db: db,
	}
}
