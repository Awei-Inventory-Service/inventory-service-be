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
