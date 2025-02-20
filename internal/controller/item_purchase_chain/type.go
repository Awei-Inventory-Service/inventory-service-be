package itempurchasechain

import (
	"github.com/gin-gonic/gin"
	itempurchasechain "github.com/inventory-service/internal/service/item_purchase_chain"
)

type ItemPurchaseChainController interface {
	Create(ctx *gin.Context)
}

type itemPurchaseChainController struct {
	itemPurchaseChainService itempurchasechain.ItemPurchaseChainService
}
