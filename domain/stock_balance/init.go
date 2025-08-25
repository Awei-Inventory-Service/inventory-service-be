package stockbalance

import (
	stockbalance "github.com/inventory-service/resource/stock_balance"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewStockBalanceDomain(
	stockBalanceResource stockbalance.StockBalanceResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
) StockBalanceDomain {
	return &stockBalanceDomain{
		stockBalanceResource:     stockBalanceResource,
		stockTransactionResource: stockTransactionResource,
	}
}
