package production

import (
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/production"
	production_item_domain "github.com/inventory-service/domain/production_item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewProductionUsecase(
	productionDomain production.ProductionDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
	inventoryDomain inventory.InventoryDomain,
	itemDomain item.ItemDomain,
	productionItemDomain production_item_domain.ProductionItemDomain,
) ProductionUsecase {
	return &productionUsecase{
		productionDomain:       productionDomain,
		stockTransactionDomain: stockTransactionDomain,
		inventoryDomain:        inventoryDomain,
		itemDomain:             itemDomain,
		productionItemDomain:   productionItemDomain,
	}
}
