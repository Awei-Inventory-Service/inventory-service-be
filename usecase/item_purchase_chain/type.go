package itempurchasechain

import (
	"context"

	"github.com/inventory-service/domain/branch"
	"github.com/inventory-service/domain/item"
	itempurchasechain "github.com/inventory-service/domain/item_purchase_chain"
	"github.com/inventory-service/domain/purchase"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ItemPurchaseChainService interface {
	Create(ctx context.Context, itemID string, branchID string, purchaseID string) *error_wrapper.ErrorWrapper
	CalculateCost(ctx context.Context, itemID string, branchID string, quantity float64) (float64, []model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper)
}

type itemPurchaseChainService struct {
	itemPurchaseChainDomain itempurchasechain.ItemPurchaseChainDomain
	purchaseDomain          purchase.PurchaseDomain
	itemDomain              item.ItemDomain
	branchDomain            branch.BranchDomain
}
