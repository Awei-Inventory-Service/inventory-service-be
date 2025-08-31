package itemcomposition

import itemcomposition "github.com/inventory-service/resource/item_composition"

func NewItemCompositionDomain(itemCompositionResource itemcomposition.ItemCompositionResourece) ItemCompositionDomain {
	return &itemCompositionDomain{
		itemCompositionResource: itemCompositionResource,
	}
}
