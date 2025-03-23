package item

import "github.com/inventory-service/app/resource/item"

func NewItemDomain(itemResource item.ItemResource) ItemDomain {
	return &itemDomain{itemResource: itemResource}
}
