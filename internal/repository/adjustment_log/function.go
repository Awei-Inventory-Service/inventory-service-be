package adjustmentlog

import (
	"github.com/inventory-service/internal/model"
)

func (a *adjustmentLogRepository) Create(adjustment model.AdjustmentLog) error {
	result := a.db.Create(&adjustment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *adjustmentLogRepository) FindAll() ([]model.AdjustmentLog, error) {
	var logs []model.AdjustmentLog
	result := a.db.Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (a *adjustmentLogRepository) FindByID(id string) (*model.AdjustmentLog, error) {
	var log model.AdjustmentLog
	result := a.db.Where("uuid = ?", id).First(&log)
	if result.Error != nil {
		return nil, result.Error
	}

	return &log, nil
}

func (a *adjustmentLogRepository) Delete(id string) error {
	result := a.db.Where("uuid = ?", id).Delete(&model.AdjustmentLog{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
