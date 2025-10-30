package product_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/mongodb"
)

type ProductSnaspshotResource interface {
	Create(ctx context.Context, payload model.ProductSnapshot) (errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.ProductSnapshot, *error_wrapper.ErrorWrapper)
	Upsert(ctx context.Context, payload dto.CreateProductSnapshotRequest) (errW *error_wrapper.ErrorWrapper)
}

type productSnapshotResource struct {
	productSnapshotCollection mongodb.MongoDBCollectionWrapper
}
