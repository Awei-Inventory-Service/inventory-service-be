package branch_item

import "gorm.io/gorm"

func NewItemBranchResource(db *gorm.DB) BranchItemResource {
	return &branchItemResource{db: db}
}
