package item

import (
	"github.com/inventory-service/resource/item"
	itemcomposition "github.com/inventory-service/resource/item_composition"
)

func NewItemDomain(
	itemResource item.ItemResource,
	itemCompositionResource itemcomposition.ItemCompositionResourece,
) ItemDomain {
	return &itemDomain{
		itemResource:            itemResource,
		itemCompositionResource: itemCompositionResource,
	}
}
