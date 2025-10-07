package dto

import "github.com/inventory-service/model"

type CreateProductRequest struct {
	Name           string                `json:"name" binding:"required"`
	Code           string                `json:"code" binding:"required"`
	Category       string                `json:"category" binding:"required"`
	ProductType    string                `json:"product_type" binding:"required"`
	Unit           string                `json:"unit" binding:"required"`
	SellingPrice   float64               `json:"selling_price" binding:"required"`
	ProductRecipes []CreateProductRecipe `json:"product_recipes" binding:"required"`
	BranchIDs      []string              `json:"branch_ids" binding:"required"`
}

type UpdateProductRequest struct {
	Name           string                `json:"name" binding:"required"`
	Code           string                `json:"code" binding:"required"`
	Category       string                `json:"category" binding:"required"`
	ProductType    string                `json:"product_type" binding:"required"`
	Unit           string                `json:"unit" binding:"required"`
	SellingPrice   float64               `json:"selling_price" binding:"required"`
	ProductRecipes []CreateProductRecipe `json:"product_recipes" binding:"required"`
	BranchIDs      []string              `json:"branch_ids" binding:"required"`
}

type CreateProductRecipe struct {
	ItemID string  `json:"item_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Unit   string  `json:"unit" binding:"required"`
	Notes  string  `json:"notes"`
}

type UpdateProductComposition struct {
	ItemID string  `json:"item_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Unit   string  `json:"unit" binding:"required"`
	Notes  string  `json:"notes"`
}
type Ingredient struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
}

type GetIngredient struct {
	ItemID      string  `json:"item_id"`
	ItemPortion float64 `json:"item_portion"`
	ItemName    string  `json:"item_name"`
	ItemUnit    string  `json:"item_unit"`
}

type GetProductResponse struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Code         string          `json:"code"`
	Category     string          `json:"category"`
	Unit         string          `json:"unit"`
	SellingPrice float64         `json:"selling_price"`
	Ingredients  []GetIngredient `json:"ingredients"`
}

type GetProductCOGSResponse struct {
	GetProductResponse
	COGS float64 `json:"cogs"`
}

func (c CreateProductRequest) MapProductCategory() model.ProductType {
	switch c.ProductType {
	case "consignment":
		return model.ProductTypeConsignment
	case "produced":
		return model.ProductTypeProduced
	}
	return ""
}
