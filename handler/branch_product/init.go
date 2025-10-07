package branch_product

import branch_product "github.com/inventory-service/usecase/branch-product"

func NewBranchProductHanlder(
	branchProductUsecase branch_product.BranchProductUsecase,
) BranchProductHandler {
	return &branchProductHandler{
		branchProductUsecase: branchProductUsecase,
	}
}
