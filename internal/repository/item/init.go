package item

import "gorm.io/gorm"

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}
