package inventory

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/inventory"
)

type InventoryHandler interface {
	FindByBranchIdAndItemId(c *gin.Context)
	FindAll(c *gin.Context)
	SyncBalance(c *gin.Context)
	Create(c *gin.Context)
	GetList(c *gin.Context)
	Recalculate(c *gin.Context)
	GetListCurrent(c *gin.Context)
}

type inventoryHandler struct {
	inventoryUsecase inventory.InventoryUsecase
}
