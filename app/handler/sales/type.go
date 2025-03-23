package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/sales"
)

type SalesController interface {
	Create(c *gin.Context)
}

type salesController struct {
	salesService sales.SalesService
}
