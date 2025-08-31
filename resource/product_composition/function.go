package productcomposition

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productCompositionResource) Create(ctx context.Context, productComposition model.ProductComposition) (errW *error_wrapper.ErrorWrapper) {
	result := p.db.Create(&productComposition)

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
	}
	return nil
}

func(p *productCompositionResource) DeleteByProductID(ctx context.Context, productID string) (errW *error_wrapper.ErrorWrapper){
	result := p.db.WithContext(ctx).Where("product_id = ?", productID).Delete(&model.ProductComposition{})

	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error)
	}

	return nil
}