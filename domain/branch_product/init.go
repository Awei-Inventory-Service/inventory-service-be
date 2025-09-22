package branch_product

import "github.com/inventory-service/resource/branch_product"

func NewBranchProductDomain(
	branchProductResource branch_product.BranchProductResource,
) BranchProductDomain {
	return &branchProductDomain{
		branchProductResource: branchProductResource,
	}
}
