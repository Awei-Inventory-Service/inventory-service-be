package inventorystockcount

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (i *inventoryStockController) Create(ctx *gin.Context) {
	var (
		inventoryStockCount dto.CreateInventoryStockCountReqest
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	if err := ctx.ShouldBindJSON(&inventoryStockCount); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.inventoryStockService.Create(ctx, inventoryStockCount.BranchID, inventoryStockCount.Items)

	if errW != nil {
		return
	}
}

func (i *inventoryStockController) Update(ctx *gin.Context) {

	var (
		inventoryStockCount dto.UpdateInventoryStockCountRequest
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&inventoryStockCount); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = i.inventoryStockService.Update(ctx, id, inventoryStockCount.BranchID, inventoryStockCount.Items)

	if errW != nil {
		return
	}
}

func (i *inventoryStockController) FindAll(ctx *gin.Context) {
	var (
		inventoryStockCounts []model.InventoryStockCount
		errW                 *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, inventoryStockCounts, errW)
	}()

	inventoryStockCounts, errW = i.inventoryStockService.FindAll(ctx)

	if errW != nil {
		return
	}

}

func (i *inventoryStockController) FindByID(ctx *gin.Context) {
	var (
		inventoryStockCount model.InventoryStockCount
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, inventoryStockCount, errW)
	}()

	id := ctx.Param("id")

	inventoryStockCount, errW = i.inventoryStockService.FindByID(ctx, id)

	if errW != nil {
		return
	}

}

func (i *inventoryStockController) FilterByBranch(ctx *gin.Context) {

	var (
		inventoryStockCounts []model.InventoryStockCount
		errW                 *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, inventoryStockCounts, errW)
	}()
	id := ctx.Param("id")

	inventoryStockCounts, errW = i.inventoryStockService.FilterByBranch(ctx, id)

	if errW != nil {
		return
	}

}

func (i *inventoryStockController) Delete(ctx *gin.Context) {

	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	errW = i.inventoryStockService.Delete(ctx, id)

	if errW != nil {
		return
	}
}
