package stockbalance

import (
	"github.com/inventory-service/resource/item"
	stockbalance "github.com/inventory-service/resource/stock_balance"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewStockBalanceDomain(
	stockBalanceResource stockbalance.StockBalanceResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
) StockBalanceDomain {
	return &stockBalanceDomain{
		stockBalanceResource:     stockBalanceResource,
		stockTransactionResource: stockTransactionResource,
		itemResource:             itemResource,
	}
}
