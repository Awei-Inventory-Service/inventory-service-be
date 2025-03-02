package upload

import (
	"github.com/inventory-service/internal/repository/sales"
	"github.com/inventory-service/lib/error_wrapper"
)

type UploadService interface {
	ParseTransactionExcel(fileName string) *error_wrapper.ErrorWrapper
}

type uploadService struct {
	salesRepository sales.SalesRepository
}
