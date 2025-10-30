package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductSnapshot struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID string             `json:"product_id" bson:"product_id"`
	BranchID  string             `json:"branch_id" bson:"branch_id"`
	Average   float64            `json:"average" bson:"average"`
	Latest    float64            `json:"latest" bson:"latest"`
	Date      time.Time          `json:"date" bson:"date"`
	Day       int                `json:"day" bson:"day"`
	Month     int                `json:"month" bson:"month"`
	Year      int                `json:"year" bson:"year"`
	Values    []struct {
		Timestamp time.Time `json:"timestamp"`
		Value     float64   `json:"value"`
	} `json:"values"`
}
