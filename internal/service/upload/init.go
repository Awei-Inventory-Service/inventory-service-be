package upload

import "github.com/inventory-service/internal/repository/sales"

func NewUploadService(salesRepository sales.SalesRepository) UploadService {
	return &uploadService{
		salesRepository: salesRepository,
	}
}
