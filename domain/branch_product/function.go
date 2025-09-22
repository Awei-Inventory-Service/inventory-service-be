package branch_product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (b *branchProductDomain) Create(ctx context.Context, payload dto.CreateBranchProductRequest) (*model.BranchProduct, *error_wrapper.ErrorWrapper) {
	branchProduct := model.BranchProduct{
		BranchID:     payload.BranchID,
		ProductID:    payload.ProductID,
		SellingPrice: &payload.SellingPrice,
	}

	// Only set BuyPrice if it's greater than 0
	if payload.BuyPrice > 0 {
		branchProduct.BuyPrice = &payload.BuyPrice
	}

	// Only set SupplierID if it's not empty
	if payload.SupplierID != "" {
		branchProduct.SupplierID = &payload.SupplierID
	}

	return b.branchProductResource.Create(ctx, branchProduct)
}

func (b *branchProductDomain) GetByBranchIdAndProductId(ctx context.Context, branchID, productID string) (*model.BranchProduct, *error_wrapper.ErrorWrapper) {
	return b.branchProductResource.GetByBranchIdAndProductId(ctx, branchID, productID)
}
