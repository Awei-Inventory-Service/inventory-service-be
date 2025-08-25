package purchase

import (
	"github.com/inventory-service/domain/branch"
	"github.com/inventory-service/domain/item"
	itempurchasechain "github.com/inventory-service/domain/item_purchase_chain"
	"github.com/inventory-service/domain/purchase"
	stockbalance "github.com/inventory-service/domain/stock_balance"
	"github.com/inventory-service/domain/supplier"
)

func NewPurchaseService(
	purchaseDomain purchase.PurchaseDomain,
	supplierDomain supplier.SupplierDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain,
	stockBalanceDomain stockbalance.StockBalanceDomain,
) PurchaseService {
	return &purchaseService{
		purchaseDomain:          purchaseDomain,
		supplierDomain:          supplierDomain,
		branchDomain:            branchDomain,
		itemDomain:              itemDomain,
		itemPurchaseChainDomain: itemPurchaseChainDomain,
		stockBalanceDomain:      stockBalanceDomain,
	}
}
