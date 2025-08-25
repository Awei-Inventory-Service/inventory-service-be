package stockbalance

import (
	"github.com/gin-gonic/gin"
	stockbalance "github.com/inventory-service/usecase/stock_balance"
)

type StockBalanceHandler interface {
	FindByBranchIdAndItemId(c *gin.Context)
}

type stockBalanceHandler struct {
	stockBalanceUsecase stockbalance.StockBalanceUsecase
}
