package model

import "time"

type ProfitSnapshot struct {
	Year      int       `json:"year"`
	Month     int       `json:"month"`
	Amount    float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
