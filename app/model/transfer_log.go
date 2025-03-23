package model

import "time"

type TransferLog struct {
	BranchOriginId string `gorm:"type:uuid;not null"`
	BranchOrigin   Branch `gorm:"foreignKey:BranchOriginId;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	BranchDestId   string `gorm:"type:uuid;not null"`
	BranchDest     Branch `gorm:"foreignKey:BranchDestId;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID         string `gorm:"type:uuid;not null"`
	Item           Item   `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	IssuerID       string `gorm:"type:uuid;not null"`
	Issuer         User   `gorm:"foreignKey:IssuerID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Quantity       int    `gorm:"type:integer;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (TransferLog) TableName() string {
	return "transfer_logs"
}
