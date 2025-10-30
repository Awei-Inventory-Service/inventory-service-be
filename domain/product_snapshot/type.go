package product_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/resource/product_snapshot"
)

type ProductSnapshotDomain interface {
	Upsert(ctx context.Context, payload dto.CreateProductSnapshotRequest) (errW *error_wrapper.ErrorWrapper)
}

type productSnapshotDomain struct {
	productSnapshotResource product_snapshot.ProductSnaspshotResource
}
