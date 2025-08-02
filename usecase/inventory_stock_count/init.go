package inventorystockcount

import (
	"github.com/inventory-service/domain/branch"
	inventorystockcount "github.com/inventory-service/domain/inventory_stock_count"
	"github.com/inventory-service/domain/item"
)

func NewInventoryStockCountService(
	inventoryStockCountDomain inventorystockcount.InventoryStockCountDomain,
	branchDomain branch.BranchDomain,
	itemDomain item.ItemDomain,
) InventoryStockCountService {
	return &inventoryStockCountService{
		inventoryStockCountDomain: inventoryStockCountDomain,
		branchDomain:              branchDomain,
		itemDomain:                itemDomain,
	}
}
