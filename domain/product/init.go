package product

import (
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
	productrecipe "github.com/inventory-service/resource/product_recipe"
)

func NewProductDomain(
	productResource product.ProductResource,
	itemResource item.ItemResource,
	productRecipeResource productrecipe.ProductRecipeResource,
	inventoryResource inventory.InventoryResource,
) ProductDomain {
	return &productDomain{
		productResource:       productResource,
		itemResource:          itemResource,
		productRecipeResource: productRecipeResource,
		inventoryResource:     inventoryResource,
	}
}
