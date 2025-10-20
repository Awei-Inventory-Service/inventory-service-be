package inventory

import (
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/inventory_snapshot"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewBranchItemDomain(
	inventoryResource inventory.InventoryResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
	purchaseResource purchase.PurchaseResource,
	inventorySnapshotResource inventory_snapshot.InventorySnapshotResource,
) InventoryDomain {
	return &inventoryDomain{
		inventoryResource:         inventoryResource,
		stockTransactionResource:  stockTransactionResource,
		itemResource:              itemResource,
		purchaseResource:          purchaseResource,
		inventorySnapshotResource: inventorySnapshotResource,
	}
}
