package model

import "time"

type Branch struct {
	UUID            string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name            string `gorm:"type:varchar(255);not null"`
	Location        string `gorm:"type:varchar(255);not null"`
	BranchManagerID string `gorm:"type:uuid;not null"`
	BranchManager   User   `gorm:"foreignKey:BranchManagerID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Branch) TableName() string {
	return "branches"
}
