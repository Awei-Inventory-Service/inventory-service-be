package stockbalance

import stockbalance "github.com/inventory-service/usecase/stock_balance"

func NewStockBalanceHandler(stockBalanceUsecase stockbalance.StockBalanceUsecase) StockBalanceHandler {
	return &stockBalanceHandler{
		stockBalanceUsecase: stockBalanceUsecase,
	}
}
