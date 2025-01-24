package model

import "time"

// Barang pembelian
type Purchase struct {
	UUID         string   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	SupplierID   string   `gorm:"type:uuid;not null"`
	Supplier     Supplier `gorm:"foreignKey:SupplierID;references:SupplierID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	BranchID     string   `gorm:"type:uuid;not null"`
	Branch       Branch   `gorm:"foreignKey:BranchID;references:BranchID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID       string   `gorm:"type:uuid;not null"`
	Item         Item     `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Quantity     int      `gorm:"type:integer;not null"`
	PurchaseCost float64  `gorm:"type:decimal;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Purchase) TableName() string {
	return "purchases"
}
