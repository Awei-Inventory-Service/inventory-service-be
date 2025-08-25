package stockbalance

import stockbalance "github.com/inventory-service/domain/stock_balance"

func NewStockBalanceUsecase(stockBalanceDomain stockbalance.StockBalanceDomain) StockBalanceUsecase {
	return &stockBalanceUsecase{
		stockBalanceDomain: stockBalanceDomain,
	}
}
