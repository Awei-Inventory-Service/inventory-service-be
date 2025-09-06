package stockbalance

import (
	"github.com/gin-gonic/gin"
	branchitem "github.com/inventory-service/usecase/branch_item"
)

type BranchItemHandler interface {
	FindByBranchIdAndItemId(c *gin.Context)
	FindAllStockBalance(c *gin.Context)
}

type itemBranchHandler struct {
	branchItemUsecase branchitem.BranchItemUsecase
}
