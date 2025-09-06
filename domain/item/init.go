package item

import (
	"github.com/inventory-service/resource/item"
	itembranch "github.com/inventory-service/resource/item_branch"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
)

func NewItemDomain(
	itemResource item.ItemResource,
	itemCompositionResource itemcomposition.ItemCompositionResourece,
	purchaseResource purchase.PurchaseResource,
	stockBalanceResource itembranch.ItemBranchResource,
) ItemDomain {
	return &itemDomain{
		itemResource:            itemResource,
		itemCompositionResource: itemCompositionResource,
		purchaseResource:        purchaseResource,
		itemBranchResource:      stockBalanceResource,
	}
}
