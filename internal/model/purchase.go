package model

import "time"

// Barang pembelian
type Purchase struct {
	UUID         string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	SupplierID   string    `gorm:"type:uuid;not null" json:"supplier_id"`
	Supplier     Supplier  `gorm:"foreignKey:SupplierID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"supplier"`
	BranchID     string    `gorm:"type:uuid;not null" json:"branch_id"`
	Branch       Branch    `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"branch"`
	ItemID       string    `gorm:"type:uuid;not null" json:"item_id"`
	Item         Item      `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"item"`
	Quantity     int       `gorm:"type:integer;not null" json:"quantity"`
	PurchaseCost float64   `gorm:"type:decimal;not null" json:"purchase_cost"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Purchase) TableName() string {
	return "purchases"
}
