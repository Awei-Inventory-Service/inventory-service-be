package inventory

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (s *inventoryHandler) Create(c *gin.Context) {
	var (
		errW                   *error_wrapper.ErrorWrapper
		createInventoryRequest dto.CreateInventoryRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createInventoryRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	userId := c.GetHeader("user_id")

	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is missing on the header")
		return
	}

	createInventoryRequest.UserID = userId

	errW = s.inventoryUsecase.Create(c, createInventoryRequest)

	if errW != nil {
		return
	}
}

func (s *inventoryHandler) FindByBranchIdAndItemId(c *gin.Context) {
	var (
		errW                  *error_wrapper.ErrorWrapper
		itemBranch            *model.Inventory
		getStockBalanceReqest dto.GetStockBalanceRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, itemBranch, errW)
	}()

	if err := c.ShouldBindJSON(&getStockBalanceReqest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	itemBranch, errW = s.inventoryUsecase.FindByBranchIdAndItemId(getStockBalanceReqest)

	if errW != nil {
		return
	}
}

func (s *inventoryHandler) FindAllBranchItem(c *gin.Context) {
	var (
		errW         *error_wrapper.ErrorWrapper
		itemBranches []dto.GetBranchItemResponse
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, itemBranches, errW)
	}()
	itemBranches, errW = s.inventoryUsecase.FindAll()
	fmt.Println("INi item branches", itemBranches)
	if errW != nil {
		fmt.Println("Errw", errW.ActualError())
		return
	}

}

func (b *inventoryHandler) SyncBalance(c *gin.Context) {
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

	errW = b.inventoryUsecase.SyncBranchItem(c, payload)
	if errW != nil {
		return
	}
}
