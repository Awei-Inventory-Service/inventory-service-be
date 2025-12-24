package sales_product_resource

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type SalesProductResource interface {
	Create(ctx context.Context, payload model.SalesProduct) (newSalesProduct model.SalesProduct, errW *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload model.SalesProduct) (newSalesProduct model.SalesProduct, errW *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, filter model.SalesProduct) (errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (salesProducts []model.SalesProduct, errW *error_wrapper.ErrorWrapper)
}

type salesProductResource struct {
	db *gorm.DB
}
