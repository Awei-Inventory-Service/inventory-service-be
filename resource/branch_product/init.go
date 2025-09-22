package branch_product

import "gorm.io/gorm"

func NewBranchProductResource(db *gorm.DB) BranchProductResource {
	return &branchProductResource{
		db: db,
	}
}
