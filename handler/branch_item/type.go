package stockbalance

import (
	"github.com/gin-gonic/gin"
	branchitem "github.com/inventory-service/usecase/branch_item"
)

type BranchItemHandler interface {
	FindByBranchIdAndItemId(c *gin.Context)
	FindAllBranchItem(c *gin.Context)
	SyncBalance(c *gin.Context)
	Create(c *gin.Context)
}

type branchItemHandler struct {
	branchItemUsecase branchitem.BranchItemUsecase
}
