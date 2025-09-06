package stockbalance

import branchitem "github.com/inventory-service/usecase/branch_item"

func NewBranchItemHandler(branchItemUsecase branchitem.BranchItemUsecase) BranchItemHandler {
	return &itemBranchHandler{
		branchItemUsecase: branchItemUsecase,
	}
}
