package item

import (
	inventory "github.com/inventory-service/resource/inventory"
	"github.com/inventory-service/resource/item"
	"github.com/inventory-service/resource/purchase"
)

func NewItemDomain(
	itemResource item.ItemResource,
	purchaseResource purchase.PurchaseResource,
	inventoryResource inventory.InventoryResource,
) ItemDomain {
	return &itemDomain{
		itemResource:      itemResource,
		purchaseResource:  purchaseResource,
		inventoryResource: inventoryResource,
	}
}
