package adjustmentlog

import "gorm.io/gorm"

func NewAdjustmentLogRepository(db *gorm.DB) AdjustmentLogRepository {
	return &adjustmentLogRepository{db: db}
}
