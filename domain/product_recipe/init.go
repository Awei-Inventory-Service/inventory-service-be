package productrecipe

import (
	productrecipe "github.com/inventory-service/resource/product_recipe"
)

func NewProductCompositionDomain(
	productRecipe productrecipe.ProductRecipeResource,
) ProductRecipeDomain {
	return &productRecipeDomain{
		productRecipe: productRecipe,
	}
}
