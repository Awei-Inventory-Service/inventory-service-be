package productrecipe

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductRecipeResource interface {
	Create(ctx context.Context, productComposition model.ProductRecipe) *error_wrapper.ErrorWrapper
	DeleteByProductID(ctx context.Context, productID string) (errW *error_wrapper.ErrorWrapper)
}

type productRecipeResource struct {
	db *gorm.DB
}
