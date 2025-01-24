package item

import "github.com/inventory-service/internal/repository/item"

func NewItemService(itemRepository item.ItemRepository) ItemService {
	return &itemService{itemRepository: itemRepository}
}
