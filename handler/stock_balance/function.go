package stockbalance

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockBalanceHandler) FindByBranchIdAndItemId(c *gin.Context) {
	var (
		errW                  *error_wrapper.ErrorWrapper
		stockBalance          *model.StockBalance
		getStockBalanceReqest dto.GetStockBalanceRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, stockBalance, errW)
	}()

	if err := c.ShouldBindJSON(&getStockBalanceReqest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	stockBalance, errW = s.stockBalanceUsecase.FindByBranchIdAndItemId(getStockBalanceReqest)

	if errW != nil {
		return
	}
	return
}

func (s *stockBalanceHandler) FindAllStockBalance(c *gin.Context) {
	var (
		errW          *error_wrapper.ErrorWrapper
		stockBalances []model.StockBalance
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, stockBalances, errW)
	}()
	stockBalances, errW = s.stockBalanceUsecase.FindAll()

	if errW != nil {
		return
	}

	return
}
