package inventory_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventorySnapshotDomain) Upsert(ctx context.Context, payload dto.CreateInventorySnapshotRequest) (errW *error_wrapper.ErrorWrapper) {
	return i.inventorySnapshotResource.Upsert(ctx, payload)
}

func (i *inventorySnapshotDomain) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.InventorySnapshot, *error_wrapper.ErrorWrapper) {
	return i.inventorySnapshotResource.Get(ctx, filter, order, limit, offset)
}
