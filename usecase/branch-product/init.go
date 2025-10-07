package branch_product

import (
	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/product"
)

func NewBranchProductUsecase(
	branchProductDomain branch_product.BranchProductDomain,
	inventoryDomain inventory.InventoryDomain,
	productDomain product.ProductDomain,
) BranchProductUsecase {
	return &branchProductUsecase{
		branchProductDomain: branchProductDomain,
		inventoryDomain:     inventoryDomain,
		productDomain:       productDomain,
	}
}
