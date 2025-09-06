package item_branch

import (
	"github.com/inventory-service/resource/item"
	itemBranch "github.com/inventory-service/resource/item_branch"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

func NewItemBranchDomain(
	itemBranchResource itemBranch.ItemBranchResource,
	stockTransactionResource stocktransaction.StockTransactionResource,
	itemResource item.ItemResource,
) ItemBranchDomain {
	return &itemBranchDomain{
		itemBranchResource:       itemBranchResource,
		stockTransactionResource: stockTransactionResource,
		itemResource:             itemResource,
	}
}
