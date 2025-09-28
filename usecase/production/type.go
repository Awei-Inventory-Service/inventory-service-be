package production

import (
	"context"

	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/production"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ProductionUsecase interface {
	Create(ctx context.Context, payload dto.CreateProductionRequest) (*model.Production, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter model.Production) ([]dto.GetProduction, *error_wrapper.ErrorWrapper)
}

type productionUsecase struct {
	productionDomain       production.ProductionDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
	inventoryDomain        inventory.InventoryDomain
}
