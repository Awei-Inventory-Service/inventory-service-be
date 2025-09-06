package item

import (
	branchitem "github.com/inventory-service/resource/branch_item"
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
)

func NewItemDomain(
	itemResource item.ItemResource,
	itemCompositionResource itemcomposition.ItemCompositionResourece,
	purchaseResource purchase.PurchaseResource,
	branchItemResource branchitem.BranchItemResource,
) ItemDomain {
	return &itemDomain{
		itemResource:            itemResource,
		itemCompositionResource: itemCompositionResource,
		purchaseResource:        purchaseResource,
		branchItemResource:      branchItemResource,
	}
}
