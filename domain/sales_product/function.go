package sales_product_domain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesProductDomain) Create(ctx context.Context, payload model.SalesProduct) (model.SalesProduct, *error_wrapper.ErrorWrapper) {
	return s.salesProductResource.Create(ctx, payload)
}
