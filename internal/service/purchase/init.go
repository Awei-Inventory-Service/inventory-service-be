package purchase

import (
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	itempurchasechain "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/purchase"
	"github.com/inventory-service/internal/repository/supplier"
)

func NewPurchaseService(
	purchaseRepository purchase.PurchaseRepository,
	supplierRepository supplier.SupplierRepository,
	branchRepository branch.BranchRepository,
	itemRepository item.ItemRepository,
	itemPurchaseChainRepository itempurchasechain.ItemPurchaseChainRepository,
) PurchaseService {
	return &purchaseService{
		purchaseRepository:          purchaseRepository,
		supplierRepository:          supplierRepository,
		branchRepository:            branchRepository,
		itemRepository:              itemRepository,
		itemPurchaseChainRepository: itemPurchaseChainRepository,
	}
}
