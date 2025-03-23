package stocktransaction

import "gorm.io/gorm"

func NewStockTransactionResource(db *gorm.DB) StockTransactionResource {
	return &stockTransactionResource{db: db}
}
