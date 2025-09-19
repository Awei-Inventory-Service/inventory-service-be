package product

import (
	"github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/product"
	productcomposition "github.com/inventory-service/resource/product_composition"
)

func NewProductDomain(
	productResource product.ProductResource,
	itemResource item.ItemResource,
	productCompositionResource productcomposition.ProductCompositionResource,
	branchItemResource branch_item.BranchItemResource,
) ProductDomain {
	return &productDomain{
		productResource:            productResource,
		itemResource:               itemResource,
		productCompositionResource: productCompositionResource,
		branchItemResource:         branchItemResource,
	}
}
