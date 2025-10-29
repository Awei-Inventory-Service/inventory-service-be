package inventory_transfer

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/inventory_transfer"
)

type InventoryTransferHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	UpdateStatus(c *gin.Context)
	GetList(c *gin.Context)
}

type inventoryTransferHandler struct {
	inventoryTransferUsecase inventory_transfer.InventoryTransferUsecase
}
