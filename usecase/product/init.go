package product

import (
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
)

func NewProductservice(productDomain product.ProductDomain, itemDomain item.ItemDomain) ProductService {
	return &productService{productDomain: productDomain, itemDomain: itemDomain}
}
