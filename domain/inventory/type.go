package inventory

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/inventory_snapshot"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type InventoryDomain interface {
	Create(branchID, itemID string, currentStock int) (*model.Inventory, *error_wrapper.ErrorWrapper)
	FindAll() (results []dto.GetInventoryResponse, errW *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.Inventory, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.Inventory, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper)
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
	SyncCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper)
	SyncBranchItem(ctx context.Context, branchID, itemID string) (currentStock, currentPrice float64, errW *error_wrapper.ErrorWrapper)
	BulkSyncBranchItems(ctx context.Context, branchID string, itemIDs []string) *error_wrapper.ErrorWrapper
	GetPrice(ctx context.Context, date dto.CustomDate, itemID, branchID string) (price float64, errW *error_wrapper.ErrorWrapper)
}

type inventoryDomain struct {
	inventoryResource         inventory.InventoryResource
	stockTransactionResource  stocktransaction.StockTransactionResource
	itemResource              item.ItemResource
	purchaseResource          purchase.PurchaseResource
	inventorySnapshotResource inventory_snapshot.InventorySnapshotResource
}
