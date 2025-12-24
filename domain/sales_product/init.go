package sales_product_domain

import sales_product_resource "github.com/inventory-service/resource/sales_product"

func NewSalesProductDomain(
	salesProductResource sales_product_resource.SalesProductResource,
) SalesProductDomain {
	return &salesProductDomain{
		salesProductResource: salesProductResource,
	}
}
