package itemcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	itemcomposition "github.com/inventory-service/resource/item_composition"
)

type ItemCompositionDomain interface {
	Create(ctx context.Context, itemComposition model.ItemComposition) (errW *error_wrapper.ErrorWrapper)
}

type itemCompositionDomain struct {
	itemCompositionResource itemcomposition.ItemCompositionResourece
}
