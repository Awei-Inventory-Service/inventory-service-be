package itempurchasechain

import (
	"context"

	"github.com/inventory-service/app/model"
	itempurchasechain "github.com/inventory-service/app/resource/item_purchase_chain"
	"github.com/inventory-service/lib/error_wrapper"
)

type ItemPurchaseChainDomain interface {
	Create(ctx context.Context, itemID string, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper
	Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper)
	BulkUpdate(ctx context.Context, payload []model.ItemPurchaseChainGet) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper
}

type itemPurchaseChainDomain struct {
	itemPurchaseChainResource itempurchasechain.ItemPurchaseChainResource
}
