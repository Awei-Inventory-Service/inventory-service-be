package dto

type CreateProductRequest struct {
	Name                string                     `json:"name" binding:"required"`
	Code                string                     `json:"code" binding:"required"`
	Category            string                     `json:"category" binding:"required"`
	Unit                string                     `json:"unit" binding:"required"`
	SellingPrice        float64                    `json:"selling_price" binding:"required"`
	ProductCompositions []CreateProductComposition `json:"product_compositions" binding:"required"`
}

type CreateProductComposition struct {
	ItemID string  `json:"item_id" binding:"required"`
	Ratio  float64 `json:"ratio" binding:"required"`
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
	Ratio       float64 `json:"ratio"`
}

type GetProductResponse struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Ingredients []GetIngredient `json:"ingredients"`
}
