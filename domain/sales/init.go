package sales

import (
	"github.com/inventory-service/resource/branch_product"
	"github.com/inventory-service/resource/product"
	"github.com/inventory-service/resource/sales"
)

func NewSalesDomain(
	salesResource sales.SalesResource,
	productResource product.ProductResource,
	branchProductResource branch_product.BranchProductResource,
) SalesDomain {
	return &salesDomain{
		salesResource:         salesResource,
		productResource:       productResource,
		branchProductResource: branchProductResource,
	}
}
