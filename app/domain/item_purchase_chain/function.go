package itempurchasechain

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemPurchaseChainDomain) Create(ctx context.Context, itemID string, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper {
	itemPurchaseChain := model.ItemPurchaseChain{
		ItemID:   itemID,
		BranchID: branchID,
		Purchase: model.ItemPurchaseChainPurchase{
			UUID:         purchase.UUID,
			ItemId:       purchase.ItemID,
			BranchId:     purchase.BranchID,
			Quantity:     purchase.Quantity,
			PurchaseCost: purchase.PurchaseCost,
		},
		Quantity: purchase.Quantity,
		Status:   model.StatusNotUsed,
		Sales:    make([]string, 0),
	}

	return i.itemPurchaseChainResource.Create(ctx, itemPurchaseChain)
}

// Find By Item Id dan BranchId
func (i *itemPurchaseChainDomain) Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper) {
	return i.itemPurchaseChainResource.Get(ctx, payload)
}

// Update -> bisa update quantity / status

func (i *itemPurchaseChainDomain) BulkUpdate(ctx context.Context, payload []model.ItemPurchaseChainGet) *error_wrapper.ErrorWrapper {
	return i.itemPurchaseChainResource.BulkUpdate(ctx, payload)
}

func (i *itemPurchaseChainDomain) Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper {
	return i.itemPurchaseChainResource.Update(ctx, id, payload)
}
