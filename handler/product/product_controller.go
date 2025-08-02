package product

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (p *productController) Create(ctx *gin.Context) {
	var (
		product dto.CreateProductRequest
		errW    *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	if err := ctx.ShouldBindJSON(&product); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = p.productService.Create(ctx, product.Name, product.Ingredients)

	if errW != nil {
		return
	}

}

func (p *productController) FindAll(ctx *gin.Context) {
	var (
		products []model.Product
		errW     *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, products, errW)
	}()

	products, errW = p.productService.FindAll(ctx)

	if errW != nil {
		return
	}

}

func (p *productController) FindByID(ctx *gin.Context) {
	var (
		product model.Product
		errW    *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, product, errW)
	}()

	id := ctx.Param("id")

	product, errW = p.productService.FindByID(ctx, id)
	if errW != nil {
		return
	}

}

func (p *productController) Update(ctx *gin.Context) {
	var (
		updatedData dto.UpdateProductRequest
		errW        *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	var ingredients []model.Ingredient

	for _, ingredient := range updatedData.Ingredients {
		ingredients = append(ingredients, model.Ingredient{
			ItemID:   ingredient.ItemID,
			Unit:     ingredient.Unit,
			Quantity: ingredient.Quantity,
		})
	}

	errW = p.productService.Update(ctx, id, updatedData.Name, ingredients)

	if errW != nil {
		return
	}

}

func (p *productController) Delete(ctx *gin.Context) {
	var errW *error_wrapper.ErrorWrapper

	defer func() {
		response_wrapper.New(&ctx.Writer, ctx, errW == nil, nil, errW)
	}()

	id := ctx.Param("id")

	errW = p.productService.Delete(ctx, id)
	if errW != nil {
		return
	}

}
