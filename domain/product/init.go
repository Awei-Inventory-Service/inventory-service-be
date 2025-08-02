package product

import "github.com/inventory-service/resource/product"

func NewProductDomain(productResource product.ProductResource) ProductDomain {
	return &productDomain{productResource: productResource}
}
