package branch

import (
	branch "github.com/inventory-service/app/domain/branch"
	user "github.com/inventory-service/app/domain/user"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

type BranchService interface {
	Create(name, location, branchManagerId string) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Branch, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Branch, *error_wrapper.ErrorWrapper)
	Update(id, name, location, branchManagerId string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type branchService struct {
	branchDomain branch.BranchDomain
	userDomain   user.UserDomain
}
