package sales

import (
	"context"

	itempurchasechain_repository "github.com/inventory-service/domain/item_purchase_chain"
	"github.com/inventory-service/domain/product"
	"github.com/inventory-service/domain/sales"
	"github.com/inventory-service/dto"

	"github.com/inventory-service/lib/error_wrapper"
	itempurchasechain "github.com/inventory-service/usecase/item_purchase_chain"
)

type SalesService interface {
	Create(ctx context.Context, payload dto.CreateSalesRequest) *error_wrapper.ErrorWrapper
}

type salesService struct {
	salesDomain              sales.SalesDomain
	productDomain            product.ProductDomain
	itemPurchaseChainDomain  itempurchasechain_repository.ItemPurchaseChainDomain
	itemPurchaseChainService itempurchasechain.ItemPurchaseChainService
}
