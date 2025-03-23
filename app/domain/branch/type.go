package branch

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/branch"
	"github.com/inventory-service/lib/error_wrapper"
)

type BranchDomain interface {
	Create(name, location, branchManagerId string) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper)
	Update(id, name, location, branchManagerId string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type branchDomain struct {
	branchResource branch.BranchResource
}
