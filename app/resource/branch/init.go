package branch

import "gorm.io/gorm"

func NewBranchResource(db *gorm.DB) BranchResource {
	return &branchResource{db: db}
}
