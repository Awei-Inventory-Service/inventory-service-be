package dto

type ProductRecipeWithPrice struct {
	UUID   string  `json:"uuid"`
	ItemID string  `json:"item_id"`
	Cost   float64 `json:"cost"`
	Unit   string  `json:"unit"`
	Amount float64 `json:"amount"`
}
