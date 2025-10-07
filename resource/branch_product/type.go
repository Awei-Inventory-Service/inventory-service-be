package branch_product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type BranchProductResource interface {
	Create(ctx context.Context, payload model.BranchProduct) (*model.BranchProduct, *error_wrapper.ErrorWrapper)
	GetByBranchIdAndProductId(ctx context.Context, branchID, productID string) (*model.BranchProduct, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, query []dto.Filter, order []dto.Order, limit, offset int) ([]model.BranchProduct, *error_wrapper.ErrorWrapper)
}

type branchProductResource struct {
	db *gorm.DB
}
