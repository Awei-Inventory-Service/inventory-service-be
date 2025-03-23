package item

import "github.com/inventory-service/app/domain/item"

func NewItemService(itemDomain item.ItemDomain) ItemService {
	return &itemService{itemDomain: itemDomain}
}
