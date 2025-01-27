package item

import "github.com/inventory-service/internal/service/item"

func NewItemController(itemService item.ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}
