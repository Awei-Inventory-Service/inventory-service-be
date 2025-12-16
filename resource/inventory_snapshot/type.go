package inventory_snapshot

import (
	"context"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/mongodb"
)

type InventorySnapshotResource interface {
	Create(ctx context.Context, payload model.InventorySnapshot) (errW *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, snapshotID string, payload model.InventorySnapshot) (errW *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, snapshotID string) (result model.InventorySnapshot, errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.InventorySnapshot, *error_wrapper.ErrorWrapper)
	Upsert(ctx context.Context, payload dto.CreateInventorySnapshotRequest) (errW *error_wrapper.ErrorWrapper)
	GetPreviousDaySnapshot(ctx context.Context, targetTime time.Time, branchID, itemID string) (*model.InventorySnapshot, *error_wrapper.ErrorWrapper)
	GetSnapshotBasedOndDate(ctx context.Context, payload dto.GetSnapshotBasedOnDateRequest) (model.InventorySnapshot, *error_wrapper.ErrorWrapper)
}

type inventorySnapshotResource struct {
	inventorySnapshotCollection mongodb.MongoDBCollectionWrapper
}
