package upload

import "github.com/inventory-service/internal/service/upload"

func NewUploadController(uploadService upload.UploadService) UploadController {
	return &uploadControllter{
		uploadService: uploadService,
	}
}
