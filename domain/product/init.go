package product

import (
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
)

func NewProductDomain(productResource product.ProductResource, itemResource item.ItemResource) ProductDomain {
	return &productDomain{
		productResource: productResource,
		itemResource:    itemResource,
	}
}
