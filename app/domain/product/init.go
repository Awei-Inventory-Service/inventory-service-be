package product

import "github.com/inventory-service/app/resource/product"

func NewProductDomain(productResource product.ProductResource) ProductDomain {
	return &productDomain{productResource: productResource}
}
