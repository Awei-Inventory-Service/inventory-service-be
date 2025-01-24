package branch

import (
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/user"
)

func NewBranchService(branchRepository branch.BranchRepository, userRepository user.UserRepository) BranchService {
	return &branchService{
		branchRepository: branchRepository,
		userRepository:   userRepository,
	}
}
