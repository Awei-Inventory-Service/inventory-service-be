package stockbalance

import (
	"github.com/gin-gonic/gin"
	itembranch "github.com/inventory-service/usecase/item_branch"
)

type ItemBranchHandler interface {
	FindByBranchIdAndItemId(c *gin.Context)
	FindAllStockBalance(c *gin.Context)
}

type itemBranchHandler struct {
	itemBranchUsecase itembranch.ItemBranchUsecase
}
