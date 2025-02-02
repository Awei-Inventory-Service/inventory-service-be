package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
)

func (p *productController) Create(ctx *gin.Context) {
	var product dto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}


	err := p.productService.Create(ctx, product.Name, product.Ingredients)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Success"})
}

func (p *productController) FindAll(ctx *gin.Context) {
	products, err := p.productService.FindAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}
	product, err := p.productService.FindByID(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	var updatedData dto.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
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

	err := p.productService.Update(ctx, id, updatedData.Name, ingredients)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Successfully update data!"})
}

func (p *productController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id is required"})
		return
	}

	err := p.productService.Delete(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully delete a product!"})
}
