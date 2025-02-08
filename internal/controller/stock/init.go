package stock

import "github.com/inventory-service/internal/service/stock"

func NewStockController(stockService stock.StockService) StockController {
	return &stockController{stockService: stockService}
}
