package production

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/production"
)

type ProductionHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type productionHandler struct {
	productionUsecase production.ProductionUsecase
}
