package model

type StockBalance struct {
	BranchID     string `gorm:"type:uuid;not null"`
	Branch       Branch `gorm:"foreignKey:BranchID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	ItemID       string `gorm:"type:uuid;not null"`
	Item         Item   `gorm:"foreignKey:ItemID;references:UUID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	CurrentStock int    `gorm:"type:integer;not null"`
}

func (StockBalance) TableName() string {
	return "stock_balances"
}
