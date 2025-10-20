package inventory_snapshot

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/resource/inventory_snapshot"
)

type InventorySnapshotDomain interface {
	Upsert(ctx context.Context, payload dto.CreateInventorySnapshotRequest) (errW *error_wrapper.ErrorWrapper)
}

type inventorySnapshotDomain struct {
	inventorySnapshotResource inventory_snapshot.InventorySnapshotResource
}
