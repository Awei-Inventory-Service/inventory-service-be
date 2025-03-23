package adjustmentlog

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (a *adjustmentLogDomain) Create(adjustment model.AdjustmentLog) *error_wrapper.ErrorWrapper {
	return a.adjusmentLogResource.Create(adjustment)
}

func (a *adjustmentLogDomain) FindAll() ([]model.AdjustmentLog, *error_wrapper.ErrorWrapper) {
	return a.adjusmentLogResource.FindAll()
}

func (a *adjustmentLogDomain) FindByID(id string) (*model.AdjustmentLog, *error_wrapper.ErrorWrapper) {
	return a.adjusmentLogResource.FindByID(id)
}

func (a *adjustmentLogDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return a.adjusmentLogResource.Delete(id)
}
