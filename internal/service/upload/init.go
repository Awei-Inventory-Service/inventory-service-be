package upload

import (
	"github.com/inventory-service/internal/repository/product"
	"github.com/inventory-service/internal/repository/sales"
	sales_service "github.com/inventory-service/internal/service/sales"
)

func NewUploadService(
	salesRepository sales.SalesRepository,
	productRepository product.ProductRepository,
	salesService sales_service.SalesService,
) UploadService {
	return &uploadService{
		salesRepository:    salesRepository,
		productRespository: productRepository,
		salesService:       salesService,
	}
}
