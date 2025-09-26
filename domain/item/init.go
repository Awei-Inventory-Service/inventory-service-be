package item

import (
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
	"github.com/inventory-service/resource/purchase"
)

func NewItemDomain(
	itemResource item.ItemResource,
	itemCompositionResource itemcomposition.ItemCompositionResourece,
	purchaseResource purchase.PurchaseResource,
	inventoryResource inventory.InventoryResource,
) ItemDomain {
	return &itemDomain{
		itemResource:            itemResource,
		itemCompositionResource: itemCompositionResource,
		purchaseResource:        purchaseResource,
		inventoryResource:       inventoryResource,
	}
}
