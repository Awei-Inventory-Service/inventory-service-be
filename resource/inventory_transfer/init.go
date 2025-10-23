package inventory_transfer

import "gorm.io/gorm"

func NewInventoryTransferResource(db *gorm.DB) InventoryTransferResource {
	return &inventoryTransferResource{db: db}
}
