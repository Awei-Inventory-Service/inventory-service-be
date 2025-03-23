package transferlog

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type TransferLogResource interface {
	Create(transferLog model.TransferLog) *error_wrapper.ErrorWrapper
	FindAll() ([]model.TransferLog, *error_wrapper.ErrorWrapper)
	FindByBranch(branchId string) ([]model.TransferLog, *error_wrapper.ErrorWrapper)
	FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, *error_wrapper.ErrorWrapper)
	Delete(branchOriginId, branchDestId, itemId string) *error_wrapper.ErrorWrapper
}

type transferLogResource struct {
	db *gorm.DB
}
