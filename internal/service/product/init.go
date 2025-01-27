package product

import "github.com/inventory-service/internal/repository/product"

func NewProductservice(productRepository product.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}
