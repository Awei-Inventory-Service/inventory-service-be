package production

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/production"
)

type ProductionHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetProductionList(c *gin.Context)
	GetByID(c *gin.Context)
}

type productionHandler struct {
	productionUsecase production.ProductionUsecase
}
