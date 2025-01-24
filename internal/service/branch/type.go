package branch

import (
	"github.com/inventory-service/internal/model"
	branch "github.com/inventory-service/internal/repository/branch"
	user "github.com/inventory-service/internal/repository/user"
)

type BranchService interface {
	Create(name, location, branchManagerId string) error
	FindAll() ([]model.Branch, error)
	FindByID(id string) (*model.Branch, error)
	Update(id, name, location, branchManagerId string) error
	Delete(id string) error
}

type branchService struct {
	branchRepository branch.BranchRepository
	userRepository   user.UserRepository
}
