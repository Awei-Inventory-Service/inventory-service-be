package inventorystockcount

import (
	"context"

	"github.com/inventory-service/domain/branch"
	inventorystockcount "github.com/inventory-service/domain/inventory_stock_count"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
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
	inventoryStockCountDomain inventorystockcount.InventoryStockCountDomain
	branchDomain              branch.BranchDomain
	itemDomain                item.ItemDomain
}
