package sales

import (
	"context"

	itempurchasechain_repository "github.com/inventory-service/app/domain/item_purchase_chain"
	"github.com/inventory-service/app/domain/product"
	"github.com/inventory-service/app/domain/sales"
	"github.com/inventory-service/app/dto"

	itempurchasechain "github.com/inventory-service/app/usecase/item_purchase_chain"
	"github.com/inventory-service/lib/error_wrapper"
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
