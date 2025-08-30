package productcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productCompositionDomain) Create(ctx context.Context, payload model.ProductComposition) (errW *error_wrapper.ErrorWrapper) {
	return p.productCompositionResource.Create(ctx, payload)
}
