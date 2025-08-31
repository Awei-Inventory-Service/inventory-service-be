package item

import "github.com/inventory-service/usecase/item"

func NewItemController(itemUsecase item.ItemUsecase) ItemController {
	return &itemController{
		itemUsecase: itemUsecase,
	}
}
