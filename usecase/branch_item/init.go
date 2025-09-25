package branch_item

import (
	branchitem "github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewBranchItemUsecase(
	branchItemDomain branchitem.BranchItemDomain,
	itemDomain item.ItemDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
) BranchItemUsecase {
	return &branchItemUsecase{
		branchItemDomain:       branchItemDomain,
		itemDomain:             itemDomain,
		stockTransactionDomain: stockTransactionDomain,
	}
}
