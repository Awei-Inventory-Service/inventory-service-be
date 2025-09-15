package dto

import "github.com/inventory-service/model"

type GetBranchItemResponse struct {
	UUID         string             `json:"uuid"`
	BranchID     string             `json:"branch_id"`
	BranchName   string             `json:"branch_name"`
	ItemID       string             `json:"item_id"`
	ItemName     string             `json:"item_name"`
	ItemCategory model.ItemCategory `json:"item_category"`
	ItemUnit     string             `json:"item_unit"`
	CurrentStock float64            `json:"current_stock"`
	Price        float64            `json:"price"`
}

type SyncBalanceRequest struct {
	ItemID   string `json:"item_id"`
	BranchID string `json:"branch_id"`
}
