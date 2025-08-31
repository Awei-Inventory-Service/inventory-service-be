package item

import (
	"github.com/inventory-service/domain/item"
	itemcomposition "github.com/inventory-service/domain/item_composition"
)

func NewItemService(
	itemDomain item.ItemDomain,
	itemCompositionDomain itemcomposition.ItemCompositionDomain,
) ItemUsecase {
	return &itemUsecase{
		itemDomain:            itemDomain,
		itemCompositionDomain: itemCompositionDomain,
	}
}
