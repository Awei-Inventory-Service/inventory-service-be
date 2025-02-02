package model

import "time"

type Invoice struct {
	UUID			string 		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FileURL			string 		`gorm:"type:text;not null"`
	Amount			float64		`gorm:"type:decimal;not null"`
	AmountOwed		float64 	`gorm:"type:decimal;not null;default:0.00"`
	Notes			string 		`gorm:"type:text"`
	Status			string 		`gorm:"type:enum('paid', 'unpaid');not null;default:'unpaid'"`
	InvoiceDate		string 		`gorm:"type:date;not null"`
	CreatedAt    	time.Time	`gorm:"autoCreateTime"`
	UpdatedAt    	time.Time	`gorm:"autoUpdateTime"`
}

func (Invoice) TableName() string {
	return "invoices"
}