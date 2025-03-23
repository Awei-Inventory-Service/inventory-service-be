package model

import "time"

type Item struct {
	UUID       string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string  `gorm:"type:varchar(255);not null"`
	Category   string  `gorm:"type:varchar(255);not null"`
	Price      float64 `gorm:"type:decimal;not null"`
	Unit       string  `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	SupplierID string   `gorm:"type:uuid;not null"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID; references:UUID;constrain:onUpdate:CASCADE,onDelete:SET NULL"`
}

func (Item) TableName() string {
	return "items"
}
