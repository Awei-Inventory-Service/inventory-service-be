package purchase

import (
	"github.com/inventory-service/app/domain/branch"
	"github.com/inventory-service/app/domain/item"
	itempurchasechain "github.com/inventory-service/app/domain/item_purchase_chain"
	"github.com/inventory-service/app/domain/purchase"
	"github.com/inventory-service/app/domain/supplier"
)

func NewPurchaseService(
	purchaseDomain purchase.PurchaseDomain,
	supplierDomain supplier.SupplierDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain,
) PurchaseService {
	return &purchaseService{
		purchaseDomain:          purchaseDomain,
		supplierDomain:          supplierDomain,
		branchDomain:            branchDomain,
		itemDomain:              itemDomain,
		itemPurchaseChainDomain: itemPurchaseChainDomain,
	}
}
