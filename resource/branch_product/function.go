package branch_product

import (
	"context"
	"errors"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

func (b *branchProductResource) Create(ctx context.Context, payload model.BranchProduct) (*model.BranchProduct, *error_wrapper.ErrorWrapper) {
	result := b.db.Create(&payload)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &payload, nil
}

func (b *branchProductResource) GetByBranchIdAndProductId(ctx context.Context, branchID, productID string) (*model.BranchProduct, *error_wrapper.ErrorWrapper) {
	var (
		branchProduct model.BranchProduct
	)

	result := b.db.Where("branch_id = ? AND product_id = ?", branchID, productID).First(&branchProduct)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Stock balance record not found")
		}

		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &branchProduct, nil
}
