package productrecipe

import "gorm.io/gorm"

func NewProductRecipeResource(db *gorm.DB) ProductRecipeResource {
	return &productRecipeResource{db: db}
}
