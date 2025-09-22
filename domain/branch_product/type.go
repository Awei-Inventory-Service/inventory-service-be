package branch_product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/branch_product"
)

type BranchProductDomain interface {
	Create(ctx context.Context, payload dto.CreateBranchProductRequest) (*model.BranchProduct, *error_wrapper.ErrorWrapper)
	GetByBranchIdAndProductId(ctx context.Context, branchID, productID string) (*model.BranchProduct, *error_wrapper.ErrorWrapper)
}

type branchProductDomain struct {
	branchProductResource branch_product.BranchProductResource
}
