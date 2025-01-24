package transferlog

import (
	"github.com/inventory-service/internal/model"
)

func (t *transferLogRepository) Create(branchOriginId, branchDestId, itemId, issuerId string, quantity int) error {
	transferLog := model.TransferLog{
		BranchOriginId: branchOriginId,
		BranchDestId:   branchDestId,
		ItemID:         itemId,
		IssuerID:       issuerId,
		Quantity:       quantity,
	}

	result := t.db.Create(&transferLog)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *transferLogRepository) FindAll() ([]model.TransferLog, error) {
	var logs []model.TransferLog
	result := t.db.Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (t *transferLogRepository) FindByBranch(branchId string) ([]model.TransferLog, error) {
	var logs []model.TransferLog
	result := t.db.Where("branch_origin_id = ? OR branch_dest_id = ?", branchId, branchId).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func (t *transferLogRepository) FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, error) {
	var log model.TransferLog
	result := t.db.Where("branch_origin_id = ? AND branch_dest_id = ? AND item_id = ?", branchOriginId, branchDestId, itemId).First(&log)
	if result.Error != nil {
		return nil, result.Error
	}

	return &log, nil
}

func (t *transferLogRepository) Delete(branchOriginId, branchDestId, itemId string) error {
	result := t.db.Where("branch_origin_id = ? AND branch_dest_id = ? AND item_id = ?", branchOriginId, branchDestId, itemId).Delete(&model.TransferLog{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
