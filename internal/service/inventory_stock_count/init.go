package inventorystockcount

import (
	"github.com/inventory-service/internal/repository/branch"
	inventorystockcount "github.com/inventory-service/internal/repository/inventory_stock_count"
	"github.com/inventory-service/internal/repository/item"
)

func NewInventoryStockCountService(
	inventoryStockCountRepository inventorystockcount.InventoryStockCountRepository,
	branchRepository branch.BranchRepository,
	itemRepository item.ItemRepository,
) InventoryStockCountService {
	return &inventoryStockCountService{
		inventoryStockCountRepository: inventoryStockCountRepository,
		branchRepository:              branchRepository,
		itemRepository:                itemRepository,
	}
}
