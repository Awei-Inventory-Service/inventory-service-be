package production_item_domain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	productionitem "github.com/inventory-service/resource/production_item"
)

type ProductionItemDomain interface {
	Create(ctx context.Context, payload model.ProductionItem) (newProductionItem *model.ProductionItem, errW *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, payload model.ProductionItem) (errW *error_wrapper.ErrorWrapper)
}

type productionItemDomain struct {
	productionItemResource productionitem.ProductionItemResource
}
