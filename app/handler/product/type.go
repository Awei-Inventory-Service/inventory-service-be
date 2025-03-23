package product

import (
	"github.com/gin-gonic/gin"
	product "github.com/inventory-service/app/usecase/product"
)

type ProductController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productController struct {
	productService product.ProductService
}
