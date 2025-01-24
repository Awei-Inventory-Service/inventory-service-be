package branch

import "gorm.io/gorm"

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepository{db: db}
}
