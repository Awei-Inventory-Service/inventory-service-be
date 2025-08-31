package itemcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemCompositionDomain) Create(ctx context.Context, itemComposition model.ItemComposition) (errW *error_wrapper.ErrorWrapper) {
	return i.itemCompositionResource.Create(ctx, itemComposition)
}
