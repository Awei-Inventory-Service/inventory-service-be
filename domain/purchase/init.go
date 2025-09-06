package purchase

import (
	itemBranch "github.com/inventory-service/resource/item_branch"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewPurchaseDomain(
	purchaseResource purchase.PurchaseResource,
	itemBranchResource itemBranch.ItemBranchResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
) PurchaseDomain {
	return &purchaseDomain{
		purchaseResource:         purchaseResource,
		itemBranchResource:       itemBranchResource,
		stockTransactionResource: stockTransactionResource,
	}
}
