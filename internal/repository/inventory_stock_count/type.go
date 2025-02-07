package inventorystockcount

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
)

type InventoryStockCountRepository interface {
	Create(ctx context.Context, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, branchID string) *error_wrapper.ErrorWrapper
}

type inventoryStockCountRepository struct {
	inventoryStockCountCollection mongodb.MongoDBCollectionWrapper
}
