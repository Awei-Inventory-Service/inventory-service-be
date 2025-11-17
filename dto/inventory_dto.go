package dto

import (
	"time"

	"github.com/inventory-service/model"
)

type GetInventoryResponse struct {
	UUID         string             `json:"uuid"`
	BranchID     string             `json:"branch_id"`
	BranchName   string             `json:"branch_name"`
	ItemID       string             `json:"item_id"`
	ItemName     string             `json:"item_name"`
	ItemCategory model.ItemCategory `json:"item_category"`
	ItemUnit     string             `json:"item_unit"`
	CurrentStock float64            `json:"current_stock"`
	Price        float64            `json:"price"`
	CreatedAt    string             `json:"created_at"`
}

type GetInventoryPriceAndValueByDate struct {
	Price   float64 `json:"price"`
	Balance float64 `json:"balance"`
	ItemID  string  `json:"item_id"`
}

type SyncBalanceRequest struct {
	ItemID   string `json:"item_id"`
	BranchID string `json:"branch_id"`
}

type CreateInventoryRequest struct {
	BranchID string  `json:"branch_id"`
	ItemID   string  `json:"item_id"`
	UserID   string  `json:"user_id"`
	Quantity float64 `json:"quantity"`
}

type RecalculateInventoryRequest struct {
	NewTime      string     `json:"start_time"`
	PreviousTime *time.Time `json:"previous_time"`
	BranchID     string     `json:"branch_id"`
	ItemID       string     `json:"item_id"`
}
