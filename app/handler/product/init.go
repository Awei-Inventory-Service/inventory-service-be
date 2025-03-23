package product

import (
	product "github.com/inventory-service/app/usecase/product"
)

func NewProductController(productService product.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}
