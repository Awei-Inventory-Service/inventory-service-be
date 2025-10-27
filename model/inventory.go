package model

import "time"

type Inventory struct {
	UUID      string  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BranchID  string  `gorm:"type:uuid;not null"`
	Branch    Branch  `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID    string  `gorm:"type:uuid;not null"`
	Item      Item    `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Stock     float64 `gorm:"type:decimal(10,2);not null"`
	Value     float64 `gorm:"type:decimal(10,2)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Inventory) TableName() string {
	return "inventories"
}
