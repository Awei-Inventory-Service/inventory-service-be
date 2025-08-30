package productcomposition

import "gorm.io/gorm"

func NewProductCompositionResource(db *gorm.DB) ProductCompositionResource {
	return &productCompositionResource{db: db}
}
