package purchase

import (
	"github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewPurchaseDomain(
	purchaseResource purchase.PurchaseResource,
	inventoryResource inventory.InventoryResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
) PurchaseDomain {
	return &purchaseDomain{
		purchaseResource:         purchaseResource,
		inventoryResource:        inventoryResource,
		stockTransactionResource: stockTransactionResource,
		itemResource:             itemResource,
	}
}
