package adjustmentlog

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type AdjustmentLogRepository interface {
	Create(adjustment model.AdjustmentLog) *error_wrapper.ErrorWrapper
	FindAll() ([]model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	Delete(id string) *error_wrapper.ErrorWrapper
}

type adjustmentLogRepository struct {
	db *gorm.DB
}
