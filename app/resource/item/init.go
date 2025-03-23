package item

import "gorm.io/gorm"

func NewItemResource(db *gorm.DB) ItemResource {
	return &itemResource{db: db}
}
