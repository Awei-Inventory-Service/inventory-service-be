package dto

type GetBranchItemResponse struct {
	UUID         string           `json:"uuid"`
	BranchID     string           `json:"branch_id"`
	ItemID       string           `json:"item_id"`
	Item         GetItemsResponse `json:"item"`
	CurrentStock float64          `json:"current_stock"`
	Price        float64          `json:"price"`
}
