package sales

import "github.com/inventory-service/resource/sales"

func NewSalesDomain(salesResource sales.SalesResource) SalesDomain {
	return &salesDomain{salesResource: salesResource}
}
