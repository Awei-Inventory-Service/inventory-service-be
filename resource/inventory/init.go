package inventory

import "gorm.io/gorm"

func NewItemBranchResource(db *gorm.DB) InventoryResource {
	return &inventoryResource{db: db}
}
