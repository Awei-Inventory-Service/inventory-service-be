package production

import (
	"context"

	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/production"
	production_item_domain "github.com/inventory-service/domain/production_item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ProductionUsecase interface {
	Create(ctx context.Context, payload dto.CreateProductionRequest) (*model.Production, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload dto.GetListRequest) (dto.GetProductionResponse, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, payload dto.DeleteProductionRequest) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, payload dto.UpdateProductionRequest) (model.Production, *error_wrapper.ErrorWrapper)
	GetByID(ctx context.Context, id string) (production dto.GetProduction, errW *error_wrapper.ErrorWrapper)
}

type productionUsecase struct {
	productionDomain       production.ProductionDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
	inventoryDomain        inventory.InventoryDomain
	itemDomain             item.ItemDomain
	productionItemDomain   production_item_domain.ProductionItemDomain
}
