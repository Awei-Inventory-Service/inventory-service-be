package product

import (
	"gorm.io/gorm"
)

func NewProductResource(db *gorm.DB) ProductResource {

	return &productResource{db: db}
}
