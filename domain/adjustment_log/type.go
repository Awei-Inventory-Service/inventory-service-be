package adjustmentlog

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	adjustmentlog "github.com/inventory-service/resource/adjustment_log"
)

type AdjusmentLogDomain interface {
	Create(adjustment model.AdjustmentLog) *error_wrapper.ErrorWrapper
	FindAll() ([]model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.AdjustmentLog, *error_wrapper.ErrorWrapper)
	Delete(id string) *error_wrapper.ErrorWrapper
}

type adjustmentLogDomain struct {
	adjusmentLogResource adjustmentlog.AdjustmentLogResource
}
