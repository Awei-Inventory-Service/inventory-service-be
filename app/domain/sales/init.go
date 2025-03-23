package sales

import "github.com/inventory-service/app/resource/sales"

func NewSalesDomain(salesResource sales.SalesResource) SalesDomain {
	return &salesDomain{salesResource: salesResource}
}
