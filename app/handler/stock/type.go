package stock

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/stock"
)

type StockController interface {
	GetStockByItemID(c *gin.Context)
}

type stockController struct {
	stockService stock.StockService
}
