package branch

import "github.com/inventory-service/usecase/branch"

func NewBranchController(branchService branch.BranchService) BranchController {
	return &branchController{
		branchService: branchService,
	}
}
