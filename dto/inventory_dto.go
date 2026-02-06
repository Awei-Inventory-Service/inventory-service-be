package dto

import (
	"time"

	"github.com/inventory-service/model"
)

type GetInventoryResponseBody struct {
	Data  []GetInventoryResponse `json:"data"`
	Count int64                  `json:"count"`
}

type GetInventoryResponse struct {
	UUID         string             `json:"uuid"`
	BranchID     string             `json:"branch_id"`
	BranchName   string             `json:"branch_name"`
	ItemID       string             `json:"item_id"`
	ItemName     string             `json:"item_name"`
	ItemCategory model.ItemCategory `json:"item_category"`
	ItemUnit     string             `json:"item_unit"`
	Stock        float64            `json:"stock"`
	Price        float64            `json:"price"`
}

type StockMovement struct {
	Date                  string  `json:"date"`
	Quantity              float64 `json:"quantity"`
	Unit                  string  `json:"unit"`
	Type                  string  `json:"type"`          // Only 2 possible values, IN / OUT
	MovementType          string  `json:"movement_type"` // It can be purchasing, sales, etc
	BranchOriginName      string  `json:"branch_origin_name"`
	BranchDestinationName string  `json:"branch_destination_name"`
}

type GetStockMovementResponse struct {
	Data  []StockMovement `json:"data"`
	Count int64           `json:"count"`
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
