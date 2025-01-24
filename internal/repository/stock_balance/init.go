package stockbalance

import "gorm.io/gorm"

func NewStockBalanceRepository(db *gorm.DB) StockBalanceRepository {
	return &stockBalanceRepository{db: db}
}
