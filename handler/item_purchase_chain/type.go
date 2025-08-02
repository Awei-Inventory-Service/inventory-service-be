package itempurchasechain

import (
	"github.com/gin-gonic/gin"
	itempurchasechain "github.com/inventory-service/usecase/item_purchase_chain"
)

type ItemPurchaseChainController interface {
	Create(ctx *gin.Context)
}

type itemPurchaseChainController struct {
	itemPurchaseChainService itempurchasechain.ItemPurchaseChainService
}
