package itempurchasechain

import itempurchasechain "github.com/inventory-service/app/resource/item_purchase_chain"

func NewItemPurchaseChainDomain(itemPurchaseChainResource itempurchasechain.ItemPurchaseChainResource) ItemPurchaseChainDomain {
	return &itemPurchaseChainDomain{itemPurchaseChainResource: itemPurchaseChainResource}
}
