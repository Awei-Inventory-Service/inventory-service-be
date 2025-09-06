package purchase

import (
	"github.com/inventory-service/domain/branch"
	"github.com/inventory-service/domain/item"
	itembranch "github.com/inventory-service/domain/item_branch"
	"github.com/inventory-service/domain/purchase"
	"github.com/inventory-service/domain/supplier"
)

func NewPurchaseService(
	purchaseDomain purchase.PurchaseDomain,
	supplierDomain supplier.SupplierDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
	stockBalanceDomain itembranch.ItemBranchDomain,
) PurchaseService {
	return &purchaseService{
		purchaseDomain:     purchaseDomain,
		supplierDomain:     supplierDomain,
		branchDomain:       branchDomain,
		itemDomain:         itemDomain,
		stockBalanceDomain: stockBalanceDomain,
	}
}
