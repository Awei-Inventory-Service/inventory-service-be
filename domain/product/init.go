package product

import (
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
	productcomposition "github.com/inventory-service/resource/product_composition"
)

func NewProductDomain(
	productResource product.ProductResource,
	itemResource item.ItemResource,
	productCompositionResource productcomposition.ProductCompositionResource,
) ProductDomain {
	return &productDomain{
		productResource:            productResource,
		itemResource:               itemResource,
		productCompositionResource: productCompositionResource,
	}
}
