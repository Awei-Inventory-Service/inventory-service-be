package itempurchasechain

import itempurchasechain "github.com/inventory-service/app/usecase/item_purchase_chain"

func NewItemPurchaseChainController(itemPurchaseChainService itempurchasechain.ItemPurchaseChainService) ItemPurchaseChainController {
	return &itemPurchaseChainController{
		itemPurchaseChainService: itemPurchaseChainService,
	}
}
