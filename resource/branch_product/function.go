package branch_product

import (
	"context"
	"errors"
	"fmt"

	"github.com/inventory-service/dto"
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

func (b *branchProductResource) Get(ctx context.Context, query []dto.Filter, order []dto.Order, limit, offset int) ([]model.BranchProduct, *error_wrapper.ErrorWrapper) {
	var (
		branchProducts []model.BranchProduct
	)

	db := b.db.Model(&model.BranchProduct{})

	// Apply filters
	for _, filter := range query {
		fmt.Println("INi len filter values", len(filter.Values))
		if len(filter.Values) == 1 {
			value := filter.Values[0]
			switch filter.Wildcard {
			case "==":
				db = db.Where(filter.Key+" = ?", value)
			case "<":
				db = db.Where(filter.Key+" < ?", value)
			}
		} else {
			db = db.Where(filter.Key+" IN ?", filter.Values)
		}
	}

	for _, ord := range order {
		if ord.IsAsc {
			db = db.Order(ord.Key + " ASC")
		} else {
			db = db.Order(ord.Key + " DESC")
		}
	}

	// Apply limit and offset
	if limit > 0 {
		db = db.Limit(limit)
	}
	if offset > 0 {
		db = db.Offset(offset)
	}

	result := db.WithContext(ctx).Preload("Branch").Preload("Product").Preload("Product.ProductRecipe").Preload("Product.ProductRecipe.Item").Find(&branchProducts)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return branchProducts, nil
}
