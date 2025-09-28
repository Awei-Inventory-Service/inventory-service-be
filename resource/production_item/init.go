package productionitem

import "gorm.io/gorm"

func NewProductionItemResource(
	db *gorm.DB,
) ProductionItemResource {
	return &productionItemResource{db: db}
}
