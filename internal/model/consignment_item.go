package model

import "time"

type ConsignmentItem struct {
	UUID      string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BranchID  string  `gorm:"type:uuid;not null"`
	Branch    Branch  `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Quantity  int     `gorm:"type:integer;not null"`
	Cost      float64 `gorm:"type:decimal;not null"`
	Price     float64 `gorm:"type:decimal;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ConsignmentItem) TableName() string {
	return "consignment_items"
}