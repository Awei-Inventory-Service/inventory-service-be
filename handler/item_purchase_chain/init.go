package itempurchasechain

import itempurchasechain "github.com/inventory-service/usecase/item_purchase_chain"

func NewItemPurchaseChainController(itemPurchaseChainService itempurchasechain.ItemPurchaseChainService) ItemPurchaseChainController {
	return &itemPurchaseChainController{
		itemPurchaseChainService: itemPurchaseChainService,
	}
}
