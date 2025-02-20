package itempurchasechain

import (
	"context"

	"github.com/inventory-service/internal/repository/branch"
	"github.com/inventory-service/internal/repository/item"
	itempurchasechain "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/purchase"
	"github.com/inventory-service/lib/error_wrapper"
)

type ItemPurchaseChainService interface {
	Create(ctx context.Context, itemID string, branchID string, purchaseID string) *error_wrapper.ErrorWrapper
}

type itemPurchaseChainService struct {
	itemPurchaseChainRepository itempurchasechain.ItemPurchaseChainRepository
	purchaseRepository          purchase.PurchaseRepository
	itemRepository              item.ItemRepository
	branchRepository            branch.BranchRepository
}
