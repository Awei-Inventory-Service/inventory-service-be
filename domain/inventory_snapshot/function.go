package inventory_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *inventorySnapshotDomain) Upsert(ctx context.Context, payload dto.CreateInventorySnapshotRequest) (errW *error_wrapper.ErrorWrapper) {
	return i.inventorySnapshotResource.Upsert(ctx, payload)
}
