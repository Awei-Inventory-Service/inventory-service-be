package model

import "time"

type InventoryStockCount struct {
	BranchID  string      `bson:"branch_id"`
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
	Items     []ItemCount `bson:"items"`
}

type ItemCount struct {
	ItemID       string `bson:"item_id"`
	ItemName     string `bson:"item_name"`
	CurrentStock int    `bson:"current_stock"`
}
