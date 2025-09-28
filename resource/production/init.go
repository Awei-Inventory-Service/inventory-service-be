package production

import "gorm.io/gorm"

func NewProductionResource(db *gorm.DB) ProductionResource {
	return &productionResource{db: db}
}
