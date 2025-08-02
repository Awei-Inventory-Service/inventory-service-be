package stock

import "github.com/inventory-service/usecase/stock"

func NewStockController(stockService stock.StockService) StockController {
	return &stockController{stockService: stockService}
}
