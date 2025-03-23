package item

import "github.com/inventory-service/app/usecase/item"

func NewItemController(itemService item.ItemService) ItemController {
	return &itemController{
		itemService: itemService,
	}
}
