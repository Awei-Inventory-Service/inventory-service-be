package item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemDomain) Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.Create(item)
}

func (i *itemDomain) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindAll()
}

func (i *itemDomain) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	return i.itemResource.FindByID(id)
}

func (i *itemDomain) Update(id, name, unit string, supplierID *string, category model.ItemCategory, price float64) *error_wrapper.ErrorWrapper {
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
