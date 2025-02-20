package itempurchasechain

import (
	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	itempurchasechain "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/purchase"
)

func NewItemPurchaseChainService(
	itemPurchaseChainRepository itempurchasechain.ItemPurchaseChainRepository,
	purchaseRepository purchase.PurchaseRepository,
	itemRepository item.ItemRepository,
	branchRepository branch.BranchRepository,
) ItemPurchaseChainService {
	return &itemPurchaseChainService{
		itemPurchaseChainRepository: itemPurchaseChainRepository,
		purchaseRepository:          purchaseRepository,
		itemRepository:              itemRepository,
		branchRepository:            branchRepository,
	}
}
