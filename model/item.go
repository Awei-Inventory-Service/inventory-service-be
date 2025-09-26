package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type ItemCategory string

// Define your item categories as constants
const (
	ItemCategoryProcessed     ItemCategory = "processed"
	ItemCategoryHalfProcessed ItemCategory = "half-processed"
	ItemCategoryRaw           ItemCategory = "raw"
	ItemCategoryOther         ItemCategory = "other"
)

// CompositionItem represents a single composition item
type CompositionItem struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	Unit     string `json:"unit"`
}

// ItemCompositions represents the JSON structure for item compositions
type ItemCompositions struct {
	Compositions []CompositionItem `json:"compositions"`
}

// Value implements the driver.Valuer interface for database storage
func (ic ItemCompositions) Value() (driver.Value, error) {
	return json.Marshal(ic)
}

// Scan implements the sql.Scanner interface for database retrieval
func (ic *ItemCompositions) Scan(value interface{}) error {
	if value == nil {
		*ic = ItemCompositions{Compositions: []CompositionItem{}}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, ic)
}

type Item struct {
	UUID         string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Name         string           `gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Category     ItemCategory     `gorm:"type:varchar(255);not null" json:"category" validate:"required"`
	Unit         string           `gorm:"type:varchar(255);not null" json:"unit" validate:"required"` // e.g., "gram", "ml"
	Compositions ItemCompositions `gorm:"type:jsonb;default:'{\"compositions\": []}'" json:"compositions"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`

	SupplierID *string  `gorm:"type:uuid" json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"supplier"`

	ProductCompositions []ProductRecipe `gorm:"foreignKey:ItemID" json:"-"`
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
