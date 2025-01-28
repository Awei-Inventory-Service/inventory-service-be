package dto

type CreateInventoryStockCountReqest struct {
	BranchID string       `json:"branch_id" binding:"required"`
	Items    []StockCount `json:"items" binding:"required"`
}
type UpdateInventoryStockCountRequest struct {
	BranchID string       `json:"branch_id" binding:"required"`
	Items    []StockCount `json:"items" binding:"required"`
}

type StockCount struct {
	ItemID       string `json:"item_id" binding:"required"`
	CurrentStock int    `json:"current_stock" binding:"required"`
}
