package stocktransaction

import "gorm.io/gorm"

func NewStockTransactionRepository(db *gorm.DB) StockTransactionRepository {
	return &stockTransactionRepository{db: db}
}
