package branch

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type BranchRepository interface {
	Create(name, location, branchManagerId string) error
	FindAll() ([]model.Branch, error)
	FindByID(id string) (*model.Branch, error)
	Update(id, name, location, branchManagerId string) error
	Delete(id string) error
}

type branchRepository struct {
	db *gorm.DB
}
