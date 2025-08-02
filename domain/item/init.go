package item

import "github.com/inventory-service/resource/item"

func NewItemDomain(itemResource item.ItemResource) ItemDomain {
	return &itemDomain{itemResource: itemResource}
}
