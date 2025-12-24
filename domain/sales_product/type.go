package sales_product_domain

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	sales_product_resource "github.com/inventory-service/resource/sales_product"
)

type SalesProductDomain interface {
	Create(ctx context.Context, payload model.SalesProduct) (salesProduct model.SalesProduct, errW *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, payload model.SalesProduct) (errW *error_wrapper.ErrorWrapper)
}

type salesProductDomain struct {
	salesProductResource sales_product_resource.SalesProductResource
}
