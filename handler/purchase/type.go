package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/usecase/purchase"
)

type PurchaseController interface {
	GetPurchases(c *gin.Context)
	GetPurchase(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type purchaseController struct {
	purchaseService purchase.PurchaseService
}
