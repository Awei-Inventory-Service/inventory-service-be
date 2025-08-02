package sales

import "github.com/inventory-service/usecase/sales"

func NewSalesController(salesService sales.SalesService) SalesController {
	return &salesController{
		salesService: salesService,
	}
}
