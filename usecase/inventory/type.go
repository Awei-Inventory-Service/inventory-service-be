package inventory

import (
	"context"

	inventory "github.com/inventory-service/domain/inventory"
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
	FindAll() ([]dto.GetBranchItemResponse, *error_wrapper.ErrorWrapper)
	SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (errW *error_wrapper.ErrorWrapper)
}

type inventoryUsecase struct {
	inventoryDomain        inventory.InventoryDomain
	itemDomain             item.ItemDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
}
