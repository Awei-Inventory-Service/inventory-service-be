package upload

import (
	"github.com/inventory-service/app/resource/product"
	"github.com/inventory-service/app/resource/sales"
	sales_service "github.com/inventory-service/app/usecase/sales"
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
