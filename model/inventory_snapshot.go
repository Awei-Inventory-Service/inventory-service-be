package model

import (
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventorySnapshot struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ItemID  string             `json:"item_id" bson:"item_id"`
	Average float64            `json:"average" bson:"average"`
	Date    time.Time          `json:"date" bson:"date"`
	Day     int                `json:"day" bson:"day"`
	Month   int                `json:"month" bson:"month"`
	Year    int                `json:"year" bson:"year"`
	Values  []struct {
		Timestamp time.Time `json:"timestamp"`
		Value     float64   `json:"value"`
	} `json:"values"`
}

func (i *InventorySnapshot) SortValuesBasedOnTimestamp() {
	sort.Slice(i.Values, func(a, b int) bool {
		return i.Values[a].Timestamp.After(i.Values[b].Timestamp)
	})
}
