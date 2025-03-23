package item

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/item"
)

type ItemController interface {
	GetItems(c *gin.Context)
	GetItem(c *gin.Context)
	CreateItem(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}

type itemController struct {
	itemService item.ItemService
}
