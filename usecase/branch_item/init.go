package branch_item

import branchitem "github.com/inventory-service/domain/branch_item"

func NewBranchItemUsecase(branchItemDomain branchitem.BranchItemDomain) BranchItemUsecase {
	return &branchItemUsecase{
		branchItemDomain: branchItemDomain,
	}
}
