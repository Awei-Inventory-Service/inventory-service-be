package item

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemDomain) Create(name, supplierId string, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	item := model.Item{
		Name:       name,
		Category:   category,
		Price:      price,
		Unit:       unit,
		SupplierID: supplierId,
	}
	return i.itemResource.Create(item)
}

func (i *itemDomain) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindAll()
}

func (i *itemDomain) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindByID(id)
}

func (i *itemDomain) Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	item := model.Item{
		Name:       name,
		Category:   category,
		Price:      price,
		Unit:       unit,
		SupplierID: supplierID,
	}

	return i.itemResource.Update(id, item)
}

func (i *itemDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return i.itemResource.Delete(id)
}
