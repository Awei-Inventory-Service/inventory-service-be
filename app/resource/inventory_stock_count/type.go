package inventorystockcount

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
)

type InventoryStockCountResource interface {
	Create(ctx context.Context, newData model.InventoryStockCount) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, stockCountId string) *error_wrapper.ErrorWrapper
}

type inventoryStockCountResource struct {
	inventoryStockCountCollection mongodb.MongoDBCollectionWrapper
}
