package sales

import (
	"context"

	"github.com/inventory-service/domain/branch_item"
	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/product"
	"github.com/inventory-service/domain/sales"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"

	"github.com/inventory-service/lib/error_wrapper"
)

type SalesService interface {
	Create(ctx context.Context, payload dto.CreateSalesRequest, userID string) *error_wrapper.ErrorWrapper
	FindGroupedByDate(ctx context.Context) ([]dto.SalesGroupedByDateResponse, *error_wrapper.ErrorWrapper)
	FindGroupedByDateAndBranch(ctx context.Context) ([]dto.SalesGroupedByDateAndBranchResponse, *error_wrapper.ErrorWrapper)
}

type salesService struct {
	salesDomain            sales.SalesDomain
	productDomain          product.ProductDomain
	branchProductDomain    branch_product.BranchProductDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
	branchItemDomain       branch_item.BranchItemDomain
}
