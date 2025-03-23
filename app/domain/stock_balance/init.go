package stockbalance

import stockbalance "github.com/inventory-service/app/resource/stock_balance"

func NewStockBalanceDomain(stockBalanceResource stockbalance.StockBalanceResource) StockBalanceDomain {
	return &stockBalanceDomain{stockBalanceResource: stockBalanceResource}
}
