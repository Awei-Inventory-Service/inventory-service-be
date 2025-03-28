package branch

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/branch"
)

type BranchController interface {
	GetBranches(c *gin.Context)
	GetBranch(c *gin.Context)
	CreateBranch(c *gin.Context)
	UpdateBranch(c *gin.Context)
	DeleteBranch(c *gin.Context)
}

type branchController struct {
	branchService branch.BranchService
}
