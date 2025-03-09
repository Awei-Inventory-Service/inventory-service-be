package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	StatusInUse   Status = "in-use"
	StatusNotUsed Status = "not-used"
	StatusUsed    Status = "used"
)

type ItemPurchaseChainPurchase struct {
	UUID         string  `json:"uuid" bson:"_id"`
	Quantity     int     `json:"quantity" bson:"quantity"`
	BranchId     string  `json:"branch_id" bson:"branch_id"`
	PurchaseCost float64 `json:"purchase_cost" bson:"purchase_cost"`
	ItemId       string  `json:"item_id" bson:"item_id"`
}

type ItemPurchaseChain struct {
	UUID     primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	ItemID   string                    `json:"item_id" bson:"item_id"`
	BranchID string                    `json:"branch_id" bson:"branch_id"`
	Purchase ItemPurchaseChainPurchase `json:"purchase" bson:"purchase"`
	Quantity int                       `json:"quantity" bson:"quantity"`
	Status   Status                    `json:"status" bson:"status"`
	Sales    []string                  `json:"sales" bson:"sales"`
}

type ItemPurchaseChainGet struct {
	ID       string                    `json:"_id"`
	ItemID   string                    `json:"item_id"`
	BranchID string                    `json:"branch_id"`
	Purchase ItemPurchaseChainPurchase `json:"purchase"`
	Quantity int                       `json:"quantity"`
	Status   Status                    `json:"status"`
	Sales    []string                  `json:"sales"`
}
