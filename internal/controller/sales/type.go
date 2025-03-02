package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/service/sales"
)

type SalesController interface {
	Create(c *gin.Context)
}

type salesController struct {
	salesService sales.SalesService
}
