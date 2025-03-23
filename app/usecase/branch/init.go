package branch

import (
	"github.com/inventory-service/app/domain/branch"
	"github.com/inventory-service/app/domain/user"
)

func NewBranchService(branchDomain branch.BranchDomain, userDomain user.UserDomain) BranchService {
	return &branchService{
		branchDomain: branchDomain,
		userDomain:   userDomain,
	}
}
