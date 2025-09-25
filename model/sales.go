package model

import "time"

type Sales struct {
	UUID            string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BranchProductID string  `gorm:"type:varchar(255);not null"`
	Type            string  `gorm:"type:varchar(255);not null"`
	Quantity        float64 `gorm:"type:decimal;not null"`
	Cost            float64 `gorm:"type:decimal;not null"`
	Price           float64   `gorm:"type:decimal;not null"`
	TransactionDate time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	BranchProduct   BranchProduct `gorm:"foreignKey:BranchProductID"`
}

func (Sales) TableName() string {
	return "sales"
}
