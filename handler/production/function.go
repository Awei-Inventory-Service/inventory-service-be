package production

import (
	"fmt"

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

func (p *productionHandler) GetProductionList(ctx *gin.Context) {
	var (
		errW                *error_wrapper.ErrorWrapper
		productionsResponse []dto.GetProduction
		filter              dto.GetListRequest
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, productionsResponse, errW)
	}()

	if err := ctx.ShouldBindJSON(&filter); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return

	}

	productionsResponse, errW = p.productionUsecase.Get(ctx, filter)
}

func (p *productionHandler) Delete(ctx *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	id := ctx.Param("id")
	userID := ctx.GetHeader("user_id")
	branchID := ctx.GetHeader("branch_id")

	if userID == "" || branchID == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id or branch id is missing the header")
		return
	}

	errW = p.productionUsecase.Delete(ctx, dto.DeleteProductionRequest{
		BranchID:     branchID,
		UserID:       userID,
		ProductionID: id,
	})
}

func (p *productionHandler) Update(ctx *gin.Context) {
	var (
		updateProductionRequest dto.UpdateProductionRequest
		errW                    *error_wrapper.ErrorWrapper
	)
	id := ctx.Param("id")

	userId := ctx.GetHeader("user_id")
	if userId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is missing on the header")
		return
	}

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	if err := ctx.ShouldBindJSON(&updateProductionRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	updateProductionRequest.ProductionID = id
	updateProductionRequest.UserID = userId
	_, errW = p.productionUsecase.Update(ctx, updateProductionRequest)
	if errW != nil {
		fmt.Println("Error updating production", errW)
		return
	}
}

func (p *productionHandler) GetByID(ctx *gin.Context) {
	var (
		productionResponse dto.GetProduction
		errW               *error_wrapper.ErrorWrapper
	)
	id := ctx.Param("id")

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, productionResponse, errW)
	}()

	productionResponse, errW = p.productionUsecase.GetByID(ctx, id)
	if errW != nil {
		fmt.Println("Error getting production by id", errW)
		return
	}
}
