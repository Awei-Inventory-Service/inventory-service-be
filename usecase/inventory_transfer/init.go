package inventory_transfer

import (
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/inventory_transfer"
	"github.com/inventory-service/domain/inventory_transfer_item"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewInventoryTransferUsecase(
	inventoryTransferDomain inventory_transfer.InventoryTransferDomain,
	inventoryTransferItemDomain inventory_transfer_item.InventoryTransferItemDomain,
	inventoryDomain inventory.InventoryDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
	itemDomain item.ItemDomain,
) InventoryTransferUsecase {
	return &inventoryTransferUsecase{
		inventoryTransferItemDomain: inventoryTransferItemDomain,
		inventoryTransferDomain:     inventoryTransferDomain,
		inventoryDomain:             inventoryDomain,
		stockTransactionDomain:      stockTransactionDomain,
		itemDomain:                  itemDomain,
	}
}
