package productcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	productcomposition "github.com/inventory-service/resource/product_composition"
)

type ProductCompositionDomain interface {
	Create(ctx context.Context, payload model.ProductComposition) (errW *error_wrapper.ErrorWrapper)
}

type productCompositionDomain struct {
	productCompositionResource productcomposition.ProductCompositionResource
}
