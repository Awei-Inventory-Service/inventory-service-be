package sales

import (
	itempurchasechain_repository "github.com/inventory-service/internal/repository/item_purchase_chain"
	"github.com/inventory-service/internal/repository/product"
	"github.com/inventory-service/internal/repository/sales"
	itempurchasechain "github.com/inventory-service/internal/service/item_purchase_chain"
)

func NewSalesService(
	salesRepository sales.SalesRepository,
	productRepository product.ProductRepository,
	itemPurchaseChainRepository itempurchasechain_repository.ItemPurchaseChainRepository,
	itemPurchaseChainService itempurchasechain.ItemPurchaseChainService,
) SalesService {
	return &salesService{
		salesRepository:             salesRepository,
		productRepository:           productRepository,
		itemPurchaseChainRepository: itemPurchaseChainRepository,
		itemPurchaseChainService:    itemPurchaseChainService,
	}
}
