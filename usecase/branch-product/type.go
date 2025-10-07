package branch_product

import (
	"context"

	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/product"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
)

type BranchProductUsecase interface {
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]dto.GetBranchProductResponse, *error_wrapper.ErrorWrapper)
}

type branchProductUsecase struct {
	branchProductDomain branch_product.BranchProductDomain
	inventoryDomain     inventory.InventoryDomain
	productDomain       product.ProductDomain
}
