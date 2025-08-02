package itempurchasechain

import (
	"github.com/inventory-service/domain/branch"
	"github.com/inventory-service/domain/item"
	itempurchasechain "github.com/inventory-service/domain/item_purchase_chain"
	"github.com/inventory-service/domain/purchase"
)

func NewItemPurchaseChainService(
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain,
	purchaseDomain purchase.PurchaseDomain,
	itemDomain item.ItemDomain,
	branchDomain branch.BranchDomain,
) ItemPurchaseChainService {
	return &itemPurchaseChainService{
		itemPurchaseChainDomain: itemPurchaseChainDomain,
		purchaseDomain:          purchaseDomain,
		itemDomain:              itemDomain,
		branchDomain:            branchDomain,
	}
}
