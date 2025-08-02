package model

import "time"

type Item struct {
	UUID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Category    string    `gorm:"type:varchar(255);not null" json:"category"`
	Price       float64   `gorm:"type:decimal;not null" json:"price"`
	Unit        string    `gorm:"type:varchar(255);not null" json:"unit"`     // e.g., "gram", "ml"
	PortionSize float64   `gorm:"type:decimal;default:1" json:"portion_size"` // default is 1 (e.g., 1 portion = 1 unit)
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	SupplierID string   `gorm:"type:uuid;not null" json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"supplier"`
}

func (Item) TableName() string {
	return "items"
}
