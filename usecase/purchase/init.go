package purchase

import (
	"github.com/inventory-service/domain/branch"
	inventory "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/purchase"
	"github.com/inventory-service/domain/supplier"
)

func NewPurchaseService(
	purchaseDomain purchase.PurchaseDomain,
	supplierDomain supplier.SupplierDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
	inventoryDomain inventory.InventoryDomain,
) PurchaseService {
	return &purchaseService{
		purchaseDomain:  purchaseDomain,
		supplierDomain:  supplierDomain,
		branchDomain:    branchDomain,
		itemDomain:      itemDomain,
		inventoryDomain: inventoryDomain,
	}
}
