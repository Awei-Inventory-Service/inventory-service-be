package adjustmentlog

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (a *adjustmentLogResource) Create(adjustment model.AdjustmentLog) *error_wrapper.ErrorWrapper {
	result := a.db.Create(&adjustment)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (a *adjustmentLogResource) FindAll() ([]model.AdjustmentLog, *error_wrapper.ErrorWrapper) {
	var logs []model.AdjustmentLog
	result := a.db.Find(&logs)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return logs, nil
}

func (a *adjustmentLogResource) FindByID(id string) (*model.AdjustmentLog, *error_wrapper.ErrorWrapper) {
	var log model.AdjustmentLog
	result := a.db.Where("uuid = ?", id).First(&log)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &log, nil
}

func (a *adjustmentLogResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := a.db.Where("uuid = ?", id).Delete(&model.AdjustmentLog{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
