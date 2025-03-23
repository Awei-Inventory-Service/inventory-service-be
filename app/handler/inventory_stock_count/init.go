package inventorystockcount

import inventorystockcount "github.com/inventory-service/app/usecase/inventory_stock_count"

func NewInventoryStockCountController(inventoryStockCountService inventorystockcount.InventoryStockCountService) InventoryStockCountController {
	return &inventoryStockController{
		inventoryStockService: inventoryStockCountService,
	}
}
