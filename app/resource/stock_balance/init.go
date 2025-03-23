package stockbalance

import "gorm.io/gorm"

func NewStockBalanceResource(db *gorm.DB) StockBalanceResource {
	return &stockBalanceResource{db: db}
}
