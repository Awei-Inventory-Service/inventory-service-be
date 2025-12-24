package sales

import (
	"github.com/inventory-service/resource/branch_product"
	"github.com/inventory-service/resource/product"
	"github.com/inventory-service/resource/sales"
	sales_product_resource "github.com/inventory-service/resource/sales_product"
)

func NewSalesDomain(
	salesResource sales.SalesResource,
	productResource product.ProductResource,
	branchProductResource branch_product.BranchProductResource,
	salesProductResource sales_product_resource.SalesProductResource,
) SalesDomain {
	return &salesDomain{
		salesResource:         salesResource,
		productResource:       productResource,
		branchProductResource: branchProductResource,
		salesProductResource:  salesProductResource,
	}
}
