package model

import (
	"time"
)

type InventoryTransfer struct {
	UUID                string                   `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BranchOriginID      string                   `gorm:"type:uuid;not null"`
	BranchOrigin        Branch                   `gorm:"foreignKey:BranchOriginID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	BranchDestinationID string                   `gorm:"type:uuid;not null"`
	BranchDestination   Branch                   `gorm:"foreignKey:BranchDestinationID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Status              string                   `gorm:"type:varchar(50);not null"`
	TransferDate        time.Time                `gorm:"not null"`
	Remarks             *string                  `gorm:"type:text"`
	IssuerID            string                   `gorm:"type:uuid;not null"`
	Items               []InventoryTransferItem  `gorm:"foreignKey:InventoryTransferID;references:UUID"`
}

type InventoryTransferItem struct {
	UUID                string            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	InventoryTransferID string            `gorm:"type:uuid;not null"`
	InventoryTransfer   InventoryTransfer `gorm:"foreignKey:InventoryTransferID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	ItemID              string            `gorm:"type:uuid;not null"`
	Item                Item              `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemQuantity        float64           `gorm:"not null"`
	ItemCost            float64           `gorm:"not null"`
	Unit                string            `gorm:"type:varchar(50)"`
}

func (InventoryTransfer) TableName() string {
	return "inventory_transfers"
}

func (InventoryTransferItem) TableName() string {
	return "inventory_transfer_items"
}
