package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/upload"
)

type UploadController interface {
	UploadTransaction(c *gin.Context)
}

type uploadControllter struct {
	uploadService upload.UploadService
}
