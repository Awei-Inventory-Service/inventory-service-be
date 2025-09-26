package productrecipe

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productRecipeResource) Create(ctx context.Context, productRecipe model.ProductRecipe) (errW *error_wrapper.ErrorWrapper) {
	result := p.db.Create(&productRecipe)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
	}
	return nil
}

func (p *productRecipeResource) DeleteByProductID(ctx context.Context, productID string) (errW *error_wrapper.ErrorWrapper) {
	result := p.db.WithContext(ctx).Where("product_id = ?", productID).Delete(&model.ProductRecipe{})

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
	}

	return nil
}
