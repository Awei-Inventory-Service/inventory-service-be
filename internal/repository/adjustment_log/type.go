package adjustmentlog

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type AdjustmentLogRepository interface {
	Create(adjustment model.AdjustmentLog) error
	FindAll() ([]model.AdjustmentLog, error)
	FindByID(id string) (*model.AdjustmentLog, error)
	Delete(id string) error
}

type adjustmentLogRepository struct {
	db *gorm.DB
}
