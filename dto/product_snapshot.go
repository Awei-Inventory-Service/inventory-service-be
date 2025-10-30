package dto

type CreateProductSnapshotRequest struct {
	ProductID string  `json:"product_id"`
	Value     float64 `json:"value"`
	BranchID  string  `json:"branch_id"`
}
