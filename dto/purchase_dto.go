package dto

type CreatePurchaseRequest struct {
	SupplierID   string  `json:"supplier_id" binding:"required"`
	BranchID     string  `json:"branch_id" binding:"required"`
	ItemID       string  `json:"item_id" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	PurchaseCost float64 `json:"purchase_cost" binding:"required"`
	Unit         string  `json:"unit" binding:"required"`
}

type UpdatePurchaseRequest struct {
	SupplierID   string  `json:"supplier_id" binding:"required"`
	BranchID     string  `json:"branch_id" binding:"required"`
	ItemID       string  `json:"item_id" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	PurchaseCost float64 `json:"purchase_cost" binding:"required"`
	Unit         string  `json:"unit" binding:"required"`
}
