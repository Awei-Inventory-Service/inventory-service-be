package production

import (
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/production"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewProductionUsecase(
	productionDomain production.ProductionDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
	inventoryDomain inventory.InventoryDomain,
) ProductionUsecase {
	return &productionUsecase{
		productionDomain:       productionDomain,
		stockTransactionDomain: stockTransactionDomain,
		inventoryDomain:        inventoryDomain,
	}
}
