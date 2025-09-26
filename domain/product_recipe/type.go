package productrecipe

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	productRecipe "github.com/inventory-service/resource/product_recipe"
)

type ProductRecipeDomain interface {
	Create(ctx context.Context, payload model.ProductRecipe) (errW *error_wrapper.ErrorWrapper)
}

type productRecipeDomain struct {
	productRecipe productRecipe.ProductRecipeResource
}
