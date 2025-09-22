package dto

type CreateSalesRequest struct {
	BranchID  string  `json:"branch_id" binding:"required"`
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required"`
	Type      string  `json:"type" binding:"required"`
}
