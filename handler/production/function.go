package production

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (p *productionHandler) Create(ctx *gin.Context) {
	var (
		createProductionRequest  dto.CreateProductionRequest
		createProductionResponse *model.Production
		errW                     *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, createProductionResponse, errW)
	}()

	if err := ctx.ShouldBindJSON(&createProductionRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	userId := ctx.GetHeader("user_id")

	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is missing on the header")
		return
	}
	createProductionRequest.UserID = userId
	createProductionResponse, errW = p.productionUsecase.Create(ctx, createProductionRequest)

}

func (p *productionHandler) Get(ctx *gin.Context) {
	var (
		errW                *error_wrapper.ErrorWrapper
		productionsResponse []dto.GetProduction
		filter              model.Production
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, productionsResponse, errW)
	}()

	if err := ctx.ShouldBindJSON(&filter); err != nil {
		return
	}

	productionsResponse, errW = p.productionUsecase.Get(ctx, filter)

}
