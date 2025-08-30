package productcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductCompositionResource interface {
	Create(ctx context.Context, productComposition model.ProductComposition) *error_wrapper.ErrorWrapper
}

type productCompositionResource struct {
	db *gorm.DB
}
