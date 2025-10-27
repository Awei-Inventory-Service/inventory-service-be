package inventory_transfer

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (h *inventoryTransferHandler) Create(ctx *gin.Context) {
	var (
		createRequest  dto.CreateInventoryTransferRequest
		createResponse model.InventoryTransfer
		errW           *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, createResponse, errW)
	}()

	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	userId := ctx.GetHeader("user_id")
	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is missing on the header")
		return
	}

	createRequest.IssuerID = userId
	createResponse, errW = h.inventoryTransferUsecase.Create(ctx, createRequest)
}

func (h *inventoryTransferHandler) UpdateStatus(ctx *gin.Context) {
	var (
		updateStatusRequest dto.UpdateInventoryTransferStatus
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	if err := ctx.ShouldBindJSON(&updateStatusRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = h.inventoryTransferUsecase.UpdateStatus(ctx, updateStatusRequest)
}

func (h *inventoryTransferHandler) GetList(ctx *gin.Context) {
	var (
		payload            dto.GetListRequest
		errW               *error_wrapper.ErrorWrapper
		inventoryTransfers dto.GetInventoryTransferListResponse
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, inventoryTransfers, errW)
	}()

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	inventoryTransfers, errW = h.inventoryTransferUsecase.Get(ctx, payload)
}
