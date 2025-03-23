package branch

import "github.com/inventory-service/app/resource/branch"

func NewBranchDomain(branchResource branch.BranchResource) BranchDomain {
	return &branchDomain{branchResource: branchResource}
}
