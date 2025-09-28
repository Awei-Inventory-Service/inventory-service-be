package item

import (
	"github.com/inventory-service/domain/item"
)

func NewItemService(
	itemDomain item.ItemDomain,
) ItemUsecase {
	return &itemUsecase{
		itemDomain: itemDomain,
	}
}
