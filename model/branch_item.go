package model

type BranchItem struct {
	UUID         string  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	BranchID     string  `gorm:"type:uuid;not null"`
	Branch       Branch  `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID       string  `gorm:"type:uuid;not null"`
	Item         Item    `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	CurrentStock float64 `gorm:"type:decimal(10,2);not null"`
	Price        float64 `gorm:"type:decimal(10,2)"`
}

func (BranchItem) TableName() string {
	return "branch_items"
}
