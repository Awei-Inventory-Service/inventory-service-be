package product_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
)

func (p *productSnapshotDomain) Upsert(ctx context.Context, payload dto.CreateProductSnapshotRequest) (errW *error_wrapper.ErrorWrapper) {
	return p.productSnapshotResource.Upsert(ctx, payload)
}
