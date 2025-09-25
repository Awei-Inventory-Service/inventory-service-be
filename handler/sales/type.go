package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/sales"
)

type SalesController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindGroupedByDate(c *gin.Context)
	FindGroupedByDateAndBranch(c *gin.Context)
}

type salesController struct {
	salesService sales.SalesService
}
