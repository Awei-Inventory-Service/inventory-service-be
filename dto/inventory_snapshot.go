package dto

import "time"

type CreateInventorySnapshotRequest struct {
	ItemID   string    `json:"item_id"`
	Value    float64   `json:"value"`
	BranchID string    `json:"branch_id"`
	Balance  float64   `json:"balance"`
	Date     time.Time `json:"date"`
}

type GetSnapshotBasedOnDateRequest struct {
	Date     time.Time
	ItemID   string
	BranchID string
}
