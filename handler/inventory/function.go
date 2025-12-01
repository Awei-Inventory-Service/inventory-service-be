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

	branchID := c.Param("branch_id")

	itemID := c.Param("item_id")
	getStockBalanceReqest.BranchId = branchID
	getStockBalanceReqest.ItemId = itemID
	itemBranch, errW = s.inventoryUsecase.FindByBranchIdAndItemId(getStockBalanceReqest)

	if errW != nil {
		return
	}
}

func (s *inventoryHandler) FindAll(c *gin.Context) {
	var (
		errW         *error_wrapper.ErrorWrapper
		itemBranches []dto.GetInventoryResponse
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

func (i *inventoryHandler) GetListCurrent(c *gin.Context) {
	var (
		errW        *error_wrapper.ErrorWrapper
		payload     dto.GetListRequest
		inventories []dto.GetInventoryResponse
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, inventories, errW)
	}()

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	branchID := c.GetHeader("branch_id")
	if branchID == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "Branch id is required in header")
		return
	}
	inventories, errW = i.inventoryUsecase.GetListCurrent(c, payload, branchID)
	if errW != nil {
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

	_, _, errW = b.inventoryUsecase.SyncBranchItem(c, payload)
	if errW != nil {
		return
	}
}
func (i *inventoryHandler) GetList(c *gin.Context) {
	var (
		errW        *error_wrapper.ErrorWrapper
		inventories []dto.GetInventoryResponse
		payload     dto.GetListRequest
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, inventories, errW)
	}()
	branchID := c.GetHeader("branch_id")
	fmt.Println("ini branch id", branchID)

	if branchID == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "Branch id is required in header")
		return
	}
	inventories, errW = i.inventoryUsecase.Get(c, payload, branchID)

	if errW != nil {
		fmt.Println("Error getting inventory", errW)
		return
	}
}

func (i *inventoryHandler) Recalculate(c *gin.Context) {
	var (
		errW    *error_wrapper.ErrorWrapper
		payload dto.RecalculateInventoryRequest
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	errW = i.inventoryUsecase.RecalculateInventory(c, payload)
}
