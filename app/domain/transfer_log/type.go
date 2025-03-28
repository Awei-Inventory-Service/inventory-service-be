package transferlog

import (
	"github.com/inventory-service/app/model"
	transferlog "github.com/inventory-service/app/resource/transfer_log"
	"github.com/inventory-service/lib/error_wrapper"
)

type TransferLogDomain interface {
	Create(branchOriginId, branchDestId, itemId, issuerId string, quantity int) *error_wrapper.ErrorWrapper
	FindAll() ([]model.TransferLog, *error_wrapper.ErrorWrapper)
	FindByBranch(branchId string) ([]model.TransferLog, *error_wrapper.ErrorWrapper)
	FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, *error_wrapper.ErrorWrapper)
	Delete(branchOriginId, branchDestId, itemId string) *error_wrapper.ErrorWrapper
}

type transferLogDomain struct {
	transferLogResource transferlog.TransferLogResource
}
