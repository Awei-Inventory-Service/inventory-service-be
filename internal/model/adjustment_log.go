package model

import "time"

type AdjustmentLog struct {
	UUID          string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BranchID      string `gorm:"type:uuid;not null"`
	Branch        Branch `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID        string `gorm:"type:uuid;not null"`
	Item          Item   `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	PreviousStock int    `gorm:"type:integer;not null"`
	NewStock      int    `gorm:"type:integer;not null"`
	AdjustorID    string `gorm:"type:uuid;not null"`
	Adjustor      User   `gorm:"foreignKey:AdjustorID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Remarks       string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (AdjustmentLog) TableName() string {
	return "adjustment_logs"
}
