package dto

type CreateProductRequest struct {
	Name        string       `json:"name" binding:"required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required,dive"`
}

type UpdateProductRequest struct {
	Name        string       `json:"name" binding:"required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required,dive"`
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
	Ratio       float64 `json:"ratio"`
}

type GetProductResponse struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Ingredients []GetIngredient `json:"ingredients"`
}
