package stock

import stockTransactionRepository "github.com/inventory-service/internal/repository/stock_transaction"

func NewStockService(stockTransactionRepository stockTransactionRepository.StockTransactionRepository) StockService {
	return &stockService{stockTransactionRepository: stockTransactionRepository}
}
