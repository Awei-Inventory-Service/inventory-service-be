package transferlog

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type TransferLogRepository interface {
	Create(branchOriginId, branchDestId, itemId, issuerId string, quantity int) error
	FindAll() ([]model.TransferLog, error)
	FindByBranch(branchId string) ([]model.TransferLog, error)
	FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, error)
	Delete(branchOriginId, branchDestId, itemId string) error
}

type transferLogRepository struct {
	db *gorm.DB
}
