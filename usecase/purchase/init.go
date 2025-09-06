package purchase

import (
	"github.com/inventory-service/domain/branch"
	branchitem "github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/purchase"
	"github.com/inventory-service/domain/supplier"
)

func NewPurchaseService(
	purchaseDomain purchase.PurchaseDomain,
	supplierDomain supplier.SupplierDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
	branchItemDomain branchitem.BranchItemDomain,
) PurchaseService {
	return &purchaseService{
		purchaseDomain:   purchaseDomain,
		supplierDomain:   supplierDomain,
		branchDomain:     branchDomain,
		itemDomain:       itemDomain,
		branchItemDomain: branchItemDomain,
	}
}
