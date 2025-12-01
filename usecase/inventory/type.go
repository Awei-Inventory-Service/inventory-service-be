package inventory

import (
	"context"

	"github.com/inventory-service/domain/branch"
	inventory "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/inventory_snapshot"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type InventoryUsecase interface {
	Create(ctx context.Context, payload dto.CreateInventoryRequest) *error_wrapper.ErrorWrapper
	FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.Inventory, *error_wrapper.ErrorWrapper)
	FindByBranchId(branchId string) ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindAll() ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper)
	SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (currentStock, currentPrice float64, errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload dto.GetListRequest, branchID string) ([]dto.GetInventoryResponse, *error_wrapper.ErrorWrapper)
	RecalculateInventory(ctx context.Context, payload dto.RecalculateInventoryRequest) (errW *error_wrapper.ErrorWrapper)
	GetListCurrent(ctx context.Context, payload dto.GetListRequest, branchID string) (inventories []dto.GetInventoryResponse, errW *error_wrapper.ErrorWrapper)
}

type inventoryUsecase struct {
	inventoryDomain         inventory.InventoryDomain
	itemDomain              item.ItemDomain
	stockTransactionDomain  stocktransaction.StockTransactionDomain
	inventorySnapshotDomain inventory_snapshot.InventorySnapshotDomain
	branchDomain            branch.BranchDomain
}
