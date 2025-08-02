package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB pakai _id, API pakai id
	Name        string             `bson:"name" json:"name"`
	Code        string             `bson:"code" json:"code"`
	Ingredients []Ingredient       `bson:"ingredients" json:"ingredients"`
}

type GetProduct struct {
	ID          string       `json:"id"`
	Code        string       `bson:"code" json:"code"`
	Name        string       `bson:"name" json:"name"`
	Ingredients []Ingredient `bson:"ingredients" json:"ingredients"`
}

type Ingredient struct {
	ID          string  `bson:"_id,omitempty" json:"id"`
	ItemID      string  `bson:"item_id" json:"item_id"`
	ItemName    string  `bson:"item_name" json:"item_name"`
	ItemPortion float64 `bson:"item_portion" json:"item_portion"`
	Quantity    int     `bson:"quantity" json:"quantity"`
	Unit        string  `bson:"unit" json:"unit"`
}
