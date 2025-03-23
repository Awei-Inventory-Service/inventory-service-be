package branch

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type BranchResource interface {
	Create(newBranch model.Branch) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper)
	Update(id string, branch model.Branch) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type branchResource struct {
	db *gorm.DB
}
