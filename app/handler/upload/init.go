package upload

import "github.com/inventory-service/app/usecase/upload"

func NewUploadController(uploadService upload.UploadService) UploadController {
	return &uploadControllter{
		uploadService: uploadService,
	}
}
