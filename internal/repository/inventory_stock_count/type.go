package inventorystockcount

import (
	"context"

	"github.com/inventory-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryStockCountRepository interface {
	Create(ctx context.Context, branchID string, items []model.ItemCount) error
	Update(ctx context.Context, stockCountID string, branchID string, items []model.ItemCount) error
	FindAll(ctx context.Context) ([]model.InventoryStockCount, error)
	FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, error)
	FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, error)
	Delete(ctx context.Context, branchID string) error
}

type inventoryStockCountRepository struct {
	inventoryStockCountCollection *mongo.Collection
}
