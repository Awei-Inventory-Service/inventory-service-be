package productrecipe

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productRecipeDomain) Create(ctx context.Context, payload model.ProductRecipe) (errW *error_wrapper.ErrorWrapper) {
	return p.productRecipe.Create(ctx, payload)
}
