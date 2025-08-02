package upload

import "github.com/inventory-service/usecase/upload"

func NewUploadController(uploadService upload.UploadService) UploadController {
	return &uploadControllter{
		uploadService: uploadService,
	}
}
