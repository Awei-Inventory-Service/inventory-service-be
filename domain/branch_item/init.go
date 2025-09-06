package branch_item

import (
	branchItem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewItemBranchDomain(
	branchItemResource branchItem.BranchItemResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
) BranchItemDomain {
	return &branchItemDomain{
		branchItemResource:       branchItemResource,
		stockTransactionResource: stockTransactionResource,
		itemResource:             itemResource,
	}
}
