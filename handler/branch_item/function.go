package stockbalance

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (s *itemBranchHandler) FindByBranchIdAndItemId(c *gin.Context) {
	var (
		errW                  *error_wrapper.ErrorWrapper
		itemBranch            *model.BranchItem
		getStockBalanceReqest dto.GetStockBalanceRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, itemBranch, errW)
	}()

	if err := c.ShouldBindJSON(&getStockBalanceReqest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	itemBranch, errW = s.branchItemUsecase.FindByBranchIdAndItemId(getStockBalanceReqest)

	if errW != nil {
		return
	}
}

func (s *itemBranchHandler) FindAllStockBalance(c *gin.Context) {
	var (
		errW         *error_wrapper.ErrorWrapper
		itemBranches []model.BranchItem
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, itemBranches, errW)
	}()
	itemBranches, errW = s.branchItemUsecase.FindAll()

	if errW != nil {
		return
	}

}
