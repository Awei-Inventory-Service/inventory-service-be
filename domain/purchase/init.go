package purchase

import (
	branchItem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewPurchaseDomain(
	purchaseResource purchase.PurchaseResource,
	branchItemResource branchItem.BranchItemResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
) PurchaseDomain {
	return &purchaseDomain{
		purchaseResource:         purchaseResource,
		branchItemResource:       branchItemResource,
		stockTransactionResource: stockTransactionResource,
	}
}
