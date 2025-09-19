package branch_item

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	branchitem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type BranchItemDomain interface {
	Create(branchID, itemID string, currentStock int) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	FindAll() (results []dto.GetBranchItemResponse, errW *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, payload model.BranchItem) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
	SyncCurrentBalance(ctx context.Context, branchID, itemID string) (float64, *error_wrapper.ErrorWrapper)
	CalculatePrice(ctx context.Context, branchID, itemID string, currentBalance float64) (float64, *error_wrapper.ErrorWrapper)
}

type branchItemDomain struct {
	branchItemResource       branchitem.BranchItemResource
	stockTransactionResource stocktransaction.StockTransactionResource
	itemResource             item.ItemResource
	purchaseResource         purchase.PurchaseResource
}
