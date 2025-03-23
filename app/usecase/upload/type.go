package upload

import (
	"context"

	"github.com/inventory-service/app/resource/product"
	"github.com/inventory-service/app/resource/sales"
	sales_service "github.com/inventory-service/app/usecase/sales"

	"github.com/inventory-service/lib/error_wrapper"
)

type UploadService interface {
	ParseTransactionExcel(ctx context.Context, fileName string, branchId string) *error_wrapper.ErrorWrapper
}

type uploadService struct {
	salesResource      sales.SalesResource
	productRespository product.ProductResource
	salesService       sales_service.SalesService
}
