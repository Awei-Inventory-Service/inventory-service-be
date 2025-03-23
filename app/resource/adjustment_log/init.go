package adjustmentlog

import "gorm.io/gorm"

func NewAdjustmentLogResource(db *gorm.DB) AdjustmentLogResource {
	return &adjustmentLogResource{db: db}
}
