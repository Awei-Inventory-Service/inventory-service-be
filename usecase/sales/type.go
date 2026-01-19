package sales

import (
	"context"

	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/domain/product"
	"github.com/inventory-service/domain/sales"
	sales_product_domain "github.com/inventory-service/domain/sales_product"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
	"github.com/inventory-service/dto"

	"github.com/inventory-service/lib/error_wrapper"
)

type SalesService interface {
	Create(ctx context.Context, payload dto.CreateSalesRequest, userID string) *error_wrapper.ErrorWrapper
	Update(ctx context.Context, payload dto.UpdateSalesRequest) (errW *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload dto.GetListRequest) ([]dto.GetSalesListResponse, *error_wrapper.ErrorWrapper)
	Delete(ctx context.Context, salesID, userID string) (errW *error_wrapper.ErrorWrapper)
}

type salesService struct {
	salesDomain            sales.SalesDomain
	productDomain          product.ProductDomain
	branchProductDomain    branch_product.BranchProductDomain
	stockTransactionDomain stocktransaction.StockTransactionDomain
	inventoryDomain        inventory.InventoryDomain
	salesProductDomain     sales_product_domain.SalesProductDomain
	itemDomain             item.ItemDomain
}
