package model

import "time"

// type: IN/OUT

type StockTransaction struct {
	UUID                string     `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BranchOriginID      string     `gorm:"type:uuid;not null"`
	BranchOrigin        Branch     `gorm:"foreignKey:BranchOriginID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID              string     `gorm:"type:uuid;not null"`
	Item                Item       `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	BranchDestinationID string     `gorm:"type:uuid;not null"`
	BranchDestination   Branch     `gorm:"foreignKey:BranchDestinationID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	IssuerID            string     `gorm:"type:uuid;not null"`
	Issuer              User       `gorm:"foreignKey:IssuerID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Type                string     `gorm:"type:varchar(255);not null"`
	Quantity            float64    `gorm:"type:decimal;not null"`
	Cost                float64    `gorm:"type:decimal;not null"`
	Reference           string     `gorm:"type:varchar(255);not null"`
	ReferenceType       *string    `gorm:"type:varchar(100)"`
	Unit                string     `gorm:"type:varchar(255);not null"`
	Remarks             string     `gorm:"type:text"`
	DeletedAt           *time.Time `gorm:"type:timestamp"`
	DeletedBy           *string    `gorm:"type:uuid"`
	DeletedByUser       *User      `gorm:"foreignKey:DeletedBy;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (StockTransaction) TableName() string {
	return "stock_transactions"
}
