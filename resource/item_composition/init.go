package itemcomposition

import (
	"gorm.io/gorm"
)

func NewItemCompositionResource(
	db *gorm.DB,
) ItemCompositionResourece {
	return &itemCompositionResource{db: db}
}
