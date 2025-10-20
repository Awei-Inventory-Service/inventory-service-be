package dto

type CreateInventorySnapshotRequest struct {
	ItemID   string  `json:"item_id"`
	Value    float64 `json:"value"`
	BranchID string  `json:"branch_id"`
}
