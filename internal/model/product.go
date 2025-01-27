package model

type Product struct {
	Name        string `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}
