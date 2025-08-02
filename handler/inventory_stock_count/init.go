package inventorystockcount

import inventorystockcount "github.com/inventory-service/usecase/inventory_stock_count"

func NewInventoryStockCountController(inventoryStockCountService inventorystockcount.InventoryStockCountService) InventoryStockCountController {
	return &inventoryStockController{
		inventoryStockService: inventoryStockCountService,
	}
}
