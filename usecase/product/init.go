package product

import (
	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
	productrecipe "github.com/inventory-service/domain/product_recipe"
)

func NewProductservice(
	productDomain product.ProductDomain,
	itemDomain item.ItemDomain,
	productRecipeDomain productrecipe.ProductRecipeDomain,
	branchProductDomain branch_product.BranchProductDomain,
) ProductService {
	return &productService{
		productDomain:       productDomain,
		itemDomain:          itemDomain,
		productRecipeDomain: productRecipeDomain,
		branchProductDomain: branchProductDomain,
	}
}
