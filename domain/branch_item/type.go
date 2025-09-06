package branch_item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	branchitem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type BranchItemDomain interface {
	Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper
	FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
	SyncCurrentBalance(branchID, itemID string) *error_wrapper.ErrorWrapper
}

type branchItemDomain struct {
	branchItemResource       branchitem.BranchItemResource
	stockTransactionResource stocktransaction.StockTransactionResource
	itemResource             item.ItemResource
}
