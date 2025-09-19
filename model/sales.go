package model

import "time"

type Sales struct {
	UUID            string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BranchProductID string  `gorm:"type:varchar(255);not null"`
	Type            string  `gorm:"type:varchar(255);not null"`
	Quantity        int     `gorm:"type:integer;not null"`
	Cost            float64 `gorm:"type:decimal;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	BranchProduct   BranchProduct `gorm:"foreignKey:BranchProductID"`
}

func (Sales) TableName() string {
	return "sales"
}
