package transferlog

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (t *transferLogDomain) Create(branchOriginId, branchDestId, itemId, issuerId string, quantity int) *error_wrapper.ErrorWrapper {
	transferLog := model.TransferLog{
		BranchOriginId: branchOriginId,
		BranchDestId:   branchDestId,
		ItemID:         itemId,
		IssuerID:       issuerId,
		Quantity:       quantity,
	}

	return t.transferLogResource.Create(transferLog)
}

func (t *transferLogDomain) FindAll() ([]model.TransferLog, *error_wrapper.ErrorWrapper) {
	return t.transferLogResource.FindAll()
}

func (t *transferLogDomain) FindByBranch(branchId string) ([]model.TransferLog, *error_wrapper.ErrorWrapper) {
	return t.transferLogResource.FindByBranch(branchId)
}

func (t *transferLogDomain) FindByID(branchOriginId, branchDestId, itemId string) (*model.TransferLog, *error_wrapper.ErrorWrapper) {
	return t.transferLogResource.FindByID(branchOriginId, branchDestId, itemId)
}

func (t *transferLogDomain) Delete(branchOriginId, branchDestId, itemId string) *error_wrapper.ErrorWrapper {
	return t.transferLogResource.Delete(branchOriginId, branchDestId, itemId)
}
