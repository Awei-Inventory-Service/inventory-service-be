package inventory

import (
	"github.com/inventory-service/domain/branch"
	inventory "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/inventory_snapshot"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewInventoryUsecase(
	inventoryDomain inventory.InventoryDomain,
	itemDomain item.ItemDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
	inventorySnapshotDomain inventory_snapshot.InventorySnapshotDomain,
	branchDomain branch.BranchDomain,
) InventoryUsecase {
	return &inventoryUsecase{
		inventoryDomain:         inventoryDomain,
		itemDomain:              itemDomain,
		stockTransactionDomain:  stockTransactionDomain,
		inventorySnapshotDomain: inventorySnapshotDomain,
		branchDomain:            branchDomain,
	}
}
