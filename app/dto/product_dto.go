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
	ItemID   string `json:"item_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
}
