package branch_item

import (
	branchItem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewBranchItemDomain(
	branchItemResource branchItem.BranchItemResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
	purchaseResource purchase.PurchaseResource,
) BranchItemDomain {
	return &branchItemDomain{
		branchItemResource:       branchItemResource,
		stockTransactionResource: stockTransactionResource,
		itemResource:             itemResource,
		purchaseResource:         purchaseResource,
	}
}
