package model

import (
	"time"
)

type Production struct {
	UUID          string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	FinalItemID   string  `gorm:"type:uuid;not null;index" json:"final_item_id" validate:"required"`
	FinalQuantity float64 `gorm:"type:decimal(10,4);not null" json:"final_quantity" validate:"required"`
	FinalUnit     string  `gorm:"type:string;not null" json:"final_unit" validate:"required"`

	BranchID       string           `gorm:"type:uuid;not null"`
	ProductionDate time.Time        `json:"production_date"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	Branch         Branch           `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	SourceItems    []ProductionItem `gorm:"foreignKey:ProductionID" json:"source_items,omitempty"`
	FinalItem      Item             `gorm:"foreignKey:FinalItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"final_item,omitempty"`
}

type ProductionItem struct {
	UUID            string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	ProductionID    string     `gorm:"type:uuid;not null;index" json:"production_id" validate:"required"`
	SourceItemID    string     `gorm:"type:uuid;not null;index" json:"source_item_id" validate:"required"`
	Quantity        float64    `gorm:"type:decimal(10,4);not null" json:"quantity" validate:"required"`
	Unit            string     `gorm:"type:string;not null" json:"unit" validate:"required"`
	WasteQuantity   float64    `gorm:"type:decimal(10,4); not null" json:"waste_quantity" validate:"required"`
	WastePercentage float64    `gorm:"type:decimal(10,4); not null" json:"waste_percentage" validate:"required"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Production      Production `gorm:"foreignKey:ProductionID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	SourceItem      Item       `gorm:"foreignKey:SourceItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"source_item,omitempty"`
}

func (ProductionItem) TableName() string {
	return "production_items"
}
func (Production) TableName() string {
	return "productions"
}
