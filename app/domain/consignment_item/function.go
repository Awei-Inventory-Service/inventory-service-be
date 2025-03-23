package consignmentitem

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (c *consignmentItemDomain) Create(branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper {
	item := model.ConsignmentItem{
		BranchID: branchID,
		Name:     name,
		Quantity: quantity,
		Cost:     cost,
		Price:    price,
	}
	return c.consignmentItemResource.Create(item)
}

func (c *consignmentItemDomain) FindAll() ([]model.ConsignmentItem, *error_wrapper.ErrorWrapper) {
	return c.consignmentItemResource.FindAll()
}

func (c *consignmentItemDomain) FindByID(id string) (*model.ConsignmentItem, *error_wrapper.ErrorWrapper) {
	return c.consignmentItemResource.FindByID(id)
}

func (c *consignmentItemDomain) Update(id, branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper {
	item := model.ConsignmentItem{
		BranchID: branchID,
		Name:     name,
		Quantity: quantity,
		Cost:     cost,
		Price:    price,
	}

	return c.consignmentItemResource.Update(id, item)
}

func (c *consignmentItemDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return c.consignmentItemResource.Delete(id)
}
