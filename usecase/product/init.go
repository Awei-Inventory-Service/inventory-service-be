package product

import (
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
	productcomposition "github.com/inventory-service/domain/product_composition"
)

func NewProductservice(
	productDomain product.ProductDomain,
	itemDomain item.ItemDomain,
	productCompositionDomain productcomposition.ProductCompositionDomain,
) ProductService {
	return &productService{
		productDomain:            productDomain,
		itemDomain:               itemDomain,
		productCompositionDomain: productCompositionDomain,
	}
}
