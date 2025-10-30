package branch_product

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
)

func (b *branchProductUsecase) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]dto.GetBranchProductResponse, *error_wrapper.ErrorWrapper) {
	var (
		results []dto.GetBranchProductResponse
	)
	branchProducts, errW := b.branchProductDomain.Get(ctx, filter, order, limit, offset)

	if errW != nil {
		return nil, errW
	}

	for _, branchProduct := range branchProducts {
		// cost := b
		_, productCost, errW := b.productDomain.CalculateProductCost(ctx, branchProduct.Product, branchProduct.BranchID, time.Now())

		if errW != nil {
			fmt.Println("Error calculating product cost", errW)
			productCost = 0
		}

		results = append(results, dto.GetBranchProductResponse{
			BranchID:        branchProduct.BranchID,
			BranchName:      branchProduct.Branch.Name,
			ProductID:       branchProduct.Product.UUID,
			ProductName:     branchProduct.Product.Name,
			ProductCategory: branchProduct.Product.Category,
			Price:           branchProduct.Product.SellingPrice,
			COGS:            productCost,
		})
	}
	return results, nil
}
