package dto

type GetStockBalanceRequest struct {
	BranchId string `json:"branch_id" binding:"required"`
	ItemId   string `json:"item_id" binding:"required"`
}
