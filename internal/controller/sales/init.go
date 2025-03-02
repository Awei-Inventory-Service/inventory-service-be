package sales

import "github.com/inventory-service/internal/service/sales"

func NewSalesController(salesService sales.SalesService) SalesController {
	return &salesController{
		salesService: salesService,
	}
}
