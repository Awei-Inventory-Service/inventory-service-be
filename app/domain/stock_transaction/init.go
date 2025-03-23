package stocktransaction

import stocktransaction "github.com/inventory-service/app/resource/stock_transaction"

func NewStockTransactionDomain(stockTransactionResource stocktransaction.StockTransactionResource) StockTransactionDomain {
	return &stockTransactionDomain{stockTransactionResource: stockTransactionResource}
}
