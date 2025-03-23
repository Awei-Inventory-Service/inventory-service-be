package itempurchasechain

import (
	"github.com/inventory-service/app/domain/branch"
	"github.com/inventory-service/app/domain/item"
	itempurchasechain "github.com/inventory-service/app/domain/item_purchase_chain"
	"github.com/inventory-service/app/domain/purchase"
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
