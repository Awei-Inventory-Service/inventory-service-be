package product

import (
	product "github.com/inventory-service/usecase/product"
)

func NewProductController(productService product.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}
