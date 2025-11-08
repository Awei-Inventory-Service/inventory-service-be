package dto

import (
	"time"

	"github.com/inventory-service/model"
)

type GetPurchaseResponse struct {
	UUID         string         `json:"uuid"`
	Supplier     model.Supplier `json:"supplier"`
	BranchID     string         `json:"branch_id"`
	BranchName   string         `json:"branch_name"`
	ItemID       string         `json:"item_id"`
	ItemName     string         `json:"item_name"`
	Quantity     float64        `json:"quantity"`
	Unit         string         `json:"unit"`
	Cost         float64        `json:"cost"`
	PurchaseDate time.Time      ` json:"purchase_date"`
}

type CreatePurchaseRequest struct {
	SupplierID   string  `json:"supplier_id" binding:"required"`
	BranchID     string  `json:"branch_id" binding:"required"`
	ItemID       string  `json:"item_id" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	PurchaseCost float64 `json:"purchase_cost" binding:"required"`
	Unit         string  `json:"unit" binding:"required"`
	PurchaseDate string  `json:"purchase_date" binding:"required"`
}

type UpdatePurchaseRequest struct {
	SupplierID   string  `json:"supplier_id" binding:"required"`
	BranchID     string  `json:"branch_id" binding:"required"`
	ItemID       string  `json:"item_id" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	PurchaseCost float64 `json:"purchase_cost" binding:"required"`
	Unit         string  `json:"unit" binding:"required"`
	UserID       string  `json:"user_id"`
	PurchaseDate string  `json:"purchase_date" binding:"required"`
}
