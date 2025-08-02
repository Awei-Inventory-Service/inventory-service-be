package branch

import "github.com/inventory-service/resource/branch"

func NewBranchDomain(branchResource branch.BranchResource) BranchDomain {
	return &branchDomain{branchResource: branchResource}
}
