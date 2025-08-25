package purchase

import (
	"github.com/inventory-service/resource/purchase"
	stockbalance "github.com/inventory-service/resource/stock_balance"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewPurchaseDomain(
	purchaseResource purchase.PurchaseResource,
	stockBalanceResource stockbalance.StockBalanceResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
) PurchaseDomain {
	return &purchaseDomain{
		purchaseResource:         purchaseResource,
		stockBalanceResource:     stockBalanceResource,
		stockTransactionResource: stockTransactionResource,
	}
}
