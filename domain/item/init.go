package item

import (
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
	stockbalance "github.com/inventory-service/resource/stock_balance"
)

func NewItemDomain(
	itemResource item.ItemResource,
	itemCompositionResource itemcomposition.ItemCompositionResourece,
	purchaseResource purchase.PurchaseResource,
	stockBalanceResource stockbalance.StockBalanceResource,
) ItemDomain {
	return &itemDomain{
		itemResource:            itemResource,
		itemCompositionResource: itemCompositionResource,
		purchaseResource:        purchaseResource,
		stockBalanceResource:    stockBalanceResource,
	}
}
