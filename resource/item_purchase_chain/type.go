package itempurchasechain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/mongodb"
)

type ItemPurchaseChainResource interface {
	Create(ctx context.Context, itemPurchaseChain model.ItemPurchaseChain) *error_wrapper.ErrorWrapper
	Get(ctx context.Context, payload model.ItemPurchaseChain) ([]model.ItemPurchaseChainGet, *error_wrapper.ErrorWrapper)
	BulkUpdate(ctx context.Context, payload []model.ItemPurchaseChainGet) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, id string, payload model.ItemPurchaseChain) *error_wrapper.ErrorWrapper
}

type itemPurchaseChainResource struct {
	itemPurchaseChainCollection mongodb.MongoDBCollectionWrapper
}
