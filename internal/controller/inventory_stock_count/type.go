package inventorystockcount

import (
	"github.com/gin-gonic/gin"
	inventorystockcount "github.com/inventory-service/internal/service/inventory_stock_count"
)

type InventoryStockCountController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FilterByBranch(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type inventoryStockController struct {
	inventoryStockService inventorystockcount.InventoryStockCountService
}
