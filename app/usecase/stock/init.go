package stock

import stockTransactionDomain "github.com/inventory-service/app/Domain/stock_transaction"

func NewStockService(stockTransactionDomain stockTransactionDomain.StockTransactionDomain) StockService {
	return &stockService{stockTransactionDomain: stockTransactionDomain}
}
