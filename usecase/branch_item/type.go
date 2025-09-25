package branch_item

import (
	"context"

	branchitem "github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/domain/item"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type BranchItemUsecase interface {
	Create(ctx context.Context, payload dto.CreateBranchItemRequest) *error_wrapper.ErrorWrapper
	FindByBranchIdAndItemId(payload dto.GetStockBalanceRequest) (*model.BranchItem, *error_wrapper.ErrorWrapper)
	FindByBranchId(branchId string) ([]model.BranchItem, *error_wrapper.ErrorWrapper)
	FindAll() ([]dto.GetBranchItemResponse, *error_wrapper.ErrorWrapper)
	SyncBranchItem(ctx context.Context, payload dto.SyncBalanceRequest) (errW *error_wrapper.ErrorWrapper)
}

type branchItemUsecase struct {
	branchItemDomain       branchitem.BranchItemDomain
	itemDomain             item.ItemDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
}
