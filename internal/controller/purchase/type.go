package purchase

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/service/purchase"
)

type PurchaseController interface {
	GetPurchases(c *gin.Context)
	GetPurchase(c *gin.Context)
	CreatePurchase(c *gin.Context)
	UpdatePurchase(c *gin.Context)
	DeletePurchase(c *gin.Context)
}

type purchaseController struct {
	purchaseService purchase.PurchaseService
}
