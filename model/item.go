package model

import "time"

type ItemCategory string

// Define your item categories as constants
const (
	ItemCategoryProcessed     ItemCategory = "processed"
	ItemCategoryHalfProcessed ItemCategory = "half-processed"
	ItemCategoryRaw           ItemCategory = "raw"
	ItemCategoryOther         ItemCategory = "other"
)

type Item struct {
	UUID        string       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name        string       `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Category    ItemCategory `gorm:"type:varchar(255);not null" json:"category" validate:"required"`
	Price       float64      `gorm:"type:decimal(10,2)" json:"price" validate:"required,gt=0"`
	Unit        string       `gorm:"type:varchar(255);not null" json:"unit" validate:"required"`       // e.g., "gram", "ml"
	PortionSize float64      `gorm:"type:decimal(10,4);default:1" json:"portion_size" validate:"gt=0"` // default is 1
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`

	SupplierID *string  `gorm:"type:uuid" json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"supplier"`

	// Relationships
	ParentCompositions  []ItemComposition    `gorm:"foreignKey:ChildItemID" json:"-"`
	ChildCompositions   []ItemComposition    `gorm:"foreignKey:ParentItemID" json:"compositions,omitempty"`
	ProductCompositions []ProductComposition `gorm:"foreignKey:ItemID" json:"-"`
}

// Fixed typo and data types
type ItemComposition struct {
	UUID         string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	ParentItemID string    `gorm:"type:uuid;not null;index" json:"parent_item_id" validate:"required"`
	ChildItemID  string    `gorm:"type:uuid;not null;index" json:"child_item_id" validate:"required"`
	Ratio        float64   `gorm:"type:decimal(10,4);not null" json:"ratio" validate:"required,gt=0"`
	Notes        string    `gorm:"type:text" json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Consider CASCADE for compositions - if parent item is deleted, composition should be too
	ParentItem Item `gorm:"foreignKey:ParentItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"parent_item,omitempty"`
	ChildItem  Item `gorm:"foreignKey:ChildItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"child_item,omitempty"`
}

func (Item) TableName() string {
	return "items"
}

func (ItemComposition) TableName() string {
	return "item_compositions"
}
