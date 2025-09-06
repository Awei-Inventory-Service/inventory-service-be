package item_branch

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/item"
	itemBranch "github.com/inventory-service/resource/item_branch"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type ItemBranchDomain interface {
	Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper
	FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.ItemBranch, *error_wrapper.ErrorWrapper)
	Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
	SyncCurrentBalance(branchID, itemID string) *error_wrapper.ErrorWrapper
}

type itemBranchDomain struct {
	itemBranchResource       itemBranch.ItemBranchResource
	stockTransactionResource stocktransaction.StockTransactionResource
	itemResource             item.ItemResource
}
