package inventorystockcount

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
)

func (i *inventoryStockController) Create(ctx *gin.Context) {
	var inventoryStockCount dto.CreateInventoryStockCountReqest

	if err := ctx.ShouldBindJSON(&inventoryStockCount); err != nil {
		ctx.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	err := i.inventoryStockService.Create(ctx, inventoryStockCount.BranchID, inventoryStockCount.Items)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully create new inventory stock count!"})
}

func (i *inventoryStockController) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	var inventoryStockCount dto.UpdateInventoryStockCountRequest

	if err := ctx.ShouldBindJSON(&inventoryStockCount); err != nil {
		ctx.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	err := i.inventoryStockService.Update(ctx, id, inventoryStockCount.BranchID, inventoryStockCount.Items)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Successfully update an inventory stock count!"})
}

func (i *inventoryStockController) FindAll(ctx *gin.Context) {
	inventoryStockCounts, err := i.inventoryStockService.FindAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inventoryStockCounts)
}

func (i *inventoryStockController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	inventoryStockCount, err := i.inventoryStockService.FindByID(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inventoryStockCount)
}

func (i *inventoryStockController) FilterByBranch(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	inventoryStockCounts, err := i.inventoryStockService.FilterByBranch(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, inventoryStockCounts)
}

func (i *inventoryStockController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	err := i.inventoryStockService.Delete(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted a stock count"})
}
