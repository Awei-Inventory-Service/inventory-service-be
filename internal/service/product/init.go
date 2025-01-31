package product

import (
	"github.com/inventory-service/internal/repository/item"
	"github.com/inventory-service/internal/repository/product"
)

func NewProductservice(productRepository product.ProductRepository, itemRepository item.ItemRepository) ProductService {
	return &productService{productRepository: productRepository, itemRepository: itemRepository}
}
