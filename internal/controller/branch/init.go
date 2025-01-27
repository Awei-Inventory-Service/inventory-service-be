package branch

import "github.com/inventory-service/internal/service/branch"

func NewBranchController(branchService branch.BranchService) BranchController {
	return &branchController{
		branchService: branchService,
	}
}
