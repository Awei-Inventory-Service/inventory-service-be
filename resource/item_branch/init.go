package item_branch

import "gorm.io/gorm"

func NewItemBranchResource(db *gorm.DB) ItemBranchResource {
	return &itemBranchResource{db: db}
}
