package transferlog

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (t *transferLogResource) Create(transferLog model.TransferLog) *error_wrapper.ErrorWrapper {
	result := t.db.Create(&transferLog)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (t *transferLogResource) FindAll() ([]model.TransferLog, *error_wrapper.ErrorWrapper) {
	var logs []model.TransferLog
	result := t.db.Find(&logs)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return logs, nil
}

func (t *transferLogResource) FindByBranch(branchId string) ([]model.TransferLog, *error_wrapper.ErrorWrapper) {
	var logs []model.TransferLog
	result := t.db.Where("branch_origin_id = ? OR branch_dest_id = ?", branchId, branchId).Find(&logs)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return logs, nil
}

func (t *transferLogResource) FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, *error_wrapper.ErrorWrapper) {
	var log model.TransferLog
	result := t.db.Where("branch_origin_id = ? AND branch_dest_id = ? AND item_id = ?", branchOriginId, branchDestId, itemId).First(&log)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &log, nil
}

func (t *transferLogResource) Delete(branchOriginId, branchDestId, itemId string) *error_wrapper.ErrorWrapper {
	result := t.db.Where("branch_origin_id = ? AND branch_dest_id = ? AND item_id = ?", branchOriginId, branchDestId, itemId).Delete(&model.TransferLog{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
