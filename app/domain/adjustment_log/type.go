package adjustmentlog

import (
	"github.com/inventory-service/app/model"
	adjustmentlog "github.com/inventory-service/app/resource/adjustment_log"
	"github.com/inventory-service/lib/error_wrapper"
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
