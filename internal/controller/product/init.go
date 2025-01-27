package product

import (
	product "github.com/inventory-service/internal/service/product"
)

func NewProductController(productService product.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}
