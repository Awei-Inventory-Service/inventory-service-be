package stock

import "github.com/inventory-service/app/usecase/stock"

func NewStockController(stockService stock.StockService) StockController {
	return &stockController{stockService: stockService}
}
