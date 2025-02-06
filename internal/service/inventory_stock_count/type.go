package inventorystockcount

import (
	"context"

	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/branch"
	inventorystockcount "github.com/inventory-service/internal/repository/inventory_stock_count"
	"github.com/inventory-service/internal/repository/item"
	"github.com/inventory-service/lib/error_wrapper"
)

type InventoryStockCountService interface {
	Create(ctx context.Context, branchID string, items []dto.StockCount) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, stockCountID string, branchID string, items []dto.StockCount) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, stockCountID string) (model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	FilterByBranch(ctx context.Context, branchID string) ([]model.InventoryStockCount, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, branchID string) *error_wrapper.ErrorWrapper
}

type inventoryStockCountService struct {
	inventoryStockCountRepository inventorystockcount.InventoryStockCountRepository
	branchRepository              branch.BranchRepository
	itemRepository                item.ItemRepository
}
