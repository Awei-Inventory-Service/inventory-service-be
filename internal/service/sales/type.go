package sales

import (
	"context"

	"github.com/inventory-service/internal/dto"
	itempurchasechain_repository "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/product"
	"github.com/inventory-service/internal/repository/sales"

	itempurchasechain "github.com/inventory-service/internal/service/item_purchase_chain"
	"github.com/inventory-service/lib/error_wrapper"
)

type SalesService interface {
	Create(ctx context.Context, payload dto.CreateSalesRequest) *error_wrapper.ErrorWrapper
}

type salesService struct {
	salesRepository             sales.SalesRepository
	productRepository           product.ProductRepository
	itemPurchaseChainRepository itempurchasechain_repository.ItemPurchaseChainRepository
	itemPurchaseChainService    itempurchasechain.ItemPurchaseChainService
}
