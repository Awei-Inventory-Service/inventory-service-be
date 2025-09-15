package stockbalance

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (s *branchItemHandler) FindByBranchIdAndItemId(c *gin.Context) {
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

func (s *branchItemHandler) FindAllBranchItem(c *gin.Context) {
	var (
		errW         *error_wrapper.ErrorWrapper
		itemBranches []dto.GetBranchItemResponse
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, itemBranches, errW)
	}()
	itemBranches, errW = s.branchItemUsecase.FindAll()
	fmt.Println("INi item branches", itemBranches)
	if errW != nil {
		fmt.Println("Errw", errW.ActualError())
		return
	}

}

func (b *branchItemHandler) SyncBalance(c *gin.Context) {
	var (
		errW    *error_wrapper.ErrorWrapper
		payload dto.SyncBalanceRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = b.branchItemUsecase.SyncBranchItem(c, payload)
	if errW != nil {
		return
	}
}
