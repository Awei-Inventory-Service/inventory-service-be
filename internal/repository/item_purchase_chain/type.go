package itempurchasechain

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
)

type ItemPurchaseChainRepository interface {
	Create(ctx context.Context, itemID string, branchID string, purchase model.Purchase) *error_wrapper.ErrorWrapper
}

type itemPurchaseChainRepository struct {
	itemPurchaseChainCollection mongodb.MongoDBCollectionWrapper
}
