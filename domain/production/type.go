package production

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/production"
	productionitem "github.com/inventory-service/resource/production_item"
)

type ProductionDomain interface {
	Create(ctx context.Context, payload dto.CreateProductionRequest) (*model.Production, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload dto.GetListRequest) ([]dto.GetProduction, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, productionID string) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, payload dto.UpdateProductionRequest) (errW *error_wrapper.ErrorWrapper)
}

type productionDomain struct {
	productionResource     production.ProductionResource
	productionItemResource productionitem.ProductionItemResource
}
