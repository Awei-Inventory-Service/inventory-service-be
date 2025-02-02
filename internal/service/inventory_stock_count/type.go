package inventorystockcount

import (
	"context"

	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/branch"
	inventorystockcount "github.com/inventory-service/internal/repository/inventory_stock_count"
	"github.com/inventory-service/internal/repository/item"
)

type InventoryStockCountService interface {
	Create(ctx context.Context, branchID string, items []dto.StockCount) error
	Update(ctx context.Context, stockCountID string, branchID string, items []dto.StockCount) error
	FindAll(ctx context.Context) ([]model.InventoryStockCount, error)
	FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, error)
	FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, error)
	Delete(ctx context.Context, branchID string) error
}

type inventoryStockCountService struct {
	inventoryStockCountRepository inventorystockcount.InventoryStockCountRepository
	branchRepository              branch.BranchRepository
	itemRepository                item.ItemRepository
}
