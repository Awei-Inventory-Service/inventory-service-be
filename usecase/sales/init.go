package sales

import (
	"github.com/inventory-service/domain/branch_product"
	"github.com/inventory-service/domain/inventory"
	"github.com/inventory-service/domain/product"
	"github.com/inventory-service/domain/sales"
	stocktransaction "github.com/inventory-service/domain/stock_transaction"
)

func NewSalesUsecase(
	salesDomain sales.SalesDomain,
	productDomain product.ProductDomain,
	branchProductDomain branch_product.BranchProductDomain,
	stockTransactionDomain stocktransaction.StockTransactionDomain,
	branchItemDomain inventory.InventoryDomain,
) SalesService {
	return &salesService{
		salesDomain:            salesDomain,
		productDomain:          productDomain,
		branchProductDomain:    branchProductDomain,
		stockTransactionDomain: stockTransactionDomain,
		branchItemDomain:       branchItemDomain,
	}
}
