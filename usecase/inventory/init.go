package inventory

import (
	inventory "github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewInventoryUsecase(
	inventoryDomain inventory.InventoryDomain,
	itemDomain item.ItemDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
) InventoryUsecase {
	return &inventoryUsecase{
		inventoryDomain:        inventoryDomain,
		itemDomain:             itemDomain,
		stockTransactionDomain: stockTransactionDomain,
	}
}
