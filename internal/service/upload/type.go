package upload

import (
	"context"

	"github.com/inventory-service/internal/repository/product"
	"github.com/inventory-service/internal/repository/sales"
	sales_service "github.com/inventory-service/internal/service/sales"

	"github.com/inventory-service/lib/error_wrapper"
)

type UploadService interface {
	ParseTransactionExcel(ctx context.Context, fileName string, branchId string) *error_wrapper.ErrorWrapper
}

type uploadService struct {
	salesRepository    sales.SalesRepository
	productRespository product.ProductRepository
	salesService       sales_service.SalesService
}
