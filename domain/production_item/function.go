package production_item_domain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productionItemDomain) Create(ctx context.Context, payload model.ProductionItem) (newProductionItem *model.ProductionItem, errW *error_wrapper.ErrorWrapper) {
	return p.productionItemResource.Create(ctx, payload)
}

func (p *productionItemDomain) Delete(ctx context.Context, payload model.ProductionItem) (errW *error_wrapper.ErrorWrapper) {
	return p.productionItemResource.Delete(ctx, payload)
}
