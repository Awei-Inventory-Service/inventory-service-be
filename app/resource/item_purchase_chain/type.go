package itempurchasechain

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
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
