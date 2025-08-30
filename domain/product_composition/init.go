package productcomposition

import productcomposition "github.com/inventory-service/resource/product_composition"

func NewProductCompositionDomain(
	productCompositionResource productcomposition.ProductCompositionResource,
) ProductCompositionDomain {
	return &productCompositionDomain{
		productCompositionResource: productCompositionResource,
	}
}
