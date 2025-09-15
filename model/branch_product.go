package model

import "time"

type BranchProduct struct {
	UUID         string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BranchID     string    `gorm:"type:uuid;not null"`
	ProductID    string    `gorm:"type:uuid;not null"`
	Stock        *float64  `gorm:"type:decimal(10,2)"`
	BuyPrice     *float64  `gorm:"type:decimal(10,2)"`
	SellingPrice *float64  `gorm:"type:decimal(10,2)"`
	SupplierID   *string   `gorm:"type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Branch       Branch    `gorm:"foreignKey:BranchID"`
	Product      Product   `gorm:"foreignKey:ProductID"`
	Supplier     *Supplier `gorm:"foreignKey:SupplierID"`
}

func (BranchProduct) TableName() string {
	return "branch_products"
}
