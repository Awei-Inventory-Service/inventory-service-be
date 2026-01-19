package model

import "time"

type Sales struct {
	UUID            string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BranchID        string         `gorm:"type:uuid;not null" json:"branch_id"`
	Branch          Branch         `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"branch"`
	TransactionDate time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	SalesProducts   []SalesProduct `gorm:"foreignKey:SalesID;references:UUID" json:"sales_products"`
}

type SalesProduct struct {
	UUID      string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	SalesID   string    `gorm:"type:uuid;not null"`
	ProductID string    `gorm:"type:uuid;not null"`
	Quantity  float64   `gorm:"type:decimal;not null"`
	Type      string    `gorm:"type:varchar(255);not null"`
	Cost      float64   `gorm:"type:decimal;not null"`
	Price     float64   `gorm:"type:decimal;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (Sales) TableName() string {
	return "sales"
}

func (SalesProduct) TableName() string {
	return "sales_products"
}
