package itemcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ItemCompositionResourece interface {
	Create(ctx context.Context, itemComposition model.ItemComposition) (errW *error_wrapper.ErrorWrapper)
}

type itemCompositionResource struct {
	db *gorm.DB
}
