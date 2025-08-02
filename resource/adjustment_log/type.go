package adjustmentlog

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type AdjustmentLogResource interface {
	Create(adjustment model.AdjustmentLog) *error_wrapper.ErrorWrapper
	FindAll() ([]model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	Delete(id string) *error_wrapper.ErrorWrapper
}

type adjustmentLogResource struct {
	db *gorm.DB
}
