package branch

import (
	"github.com/inventory-service/internal/model"
	branch "github.com/inventory-service/internal/repository/branch"
	user "github.com/inventory-service/internal/repository/user"
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
	branchRepository branch.BranchRepository
	userRepository   user.UserRepository
}
