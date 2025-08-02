package upload

import (
	"github.com/inventory-service/resource/product"
	"github.com/inventory-service/resource/sales"
	sales_service "github.com/inventory-service/usecase/sales"
)

func NewUploadService(
	salesResource sales.SalesResource,
	productResource product.ProductResource,
	salesService sales_service.SalesService,
) UploadService {
	return &uploadService{
		salesResource:      salesResource,
		productRespository: productResource,
		salesService:       salesService,
	}
}
