package sales

import (
	itempurchasechain_repository "github.com/inventory-service/app/domain/item_purchase_chain"
	"github.com/inventory-service/app/domain/product"
	"github.com/inventory-service/app/domain/sales"
	itempurchasechain "github.com/inventory-service/app/usecase/item_purchase_chain"
)

func NewSalesService(
	salesDomain sales.SalesDomain,
	productDomain product.ProductDomain,
	itemPurchaseChainDomain itempurchasechain_repository.ItemPurchaseChainDomain,
	itemPurchaseChainService itempurchasechain.ItemPurchaseChainService,
) SalesService {
	return &salesService{
		salesDomain:              salesDomain,
		productDomain:            productDomain,
		itemPurchaseChainDomain:  itemPurchaseChainDomain,
		itemPurchaseChainService: itemPurchaseChainService,
	}
}
