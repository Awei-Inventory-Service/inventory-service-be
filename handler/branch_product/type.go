package branch_product

import (
	"github.com/gin-gonic/gin"
	branch_product "github.com/inventory-service/usecase/branch-product"
)

type BranchProductHandler interface {
	GetBranchProductList(ctx *gin.Context)
}

type branchProductHandler struct {
	branchProductUsecase branch_product.BranchProductUsecase
}
