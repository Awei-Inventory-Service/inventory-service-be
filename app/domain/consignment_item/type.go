package consignmentitem

import (
	"github.com/inventory-service/app/model"
	consignmentitem "github.com/inventory-service/app/resource/consignment_item"
	"github.com/inventory-service/lib/error_wrapper"
)

type ConsignmentItemDomain interface {
	Create(branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	Update(id, branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type consignmentItemDomain struct {
	consignmentItemResource consignmentitem.ConsignmentItemResource
}
