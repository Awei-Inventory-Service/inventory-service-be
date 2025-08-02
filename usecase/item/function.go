package item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemService) Create(name, supplierID, category, unit string, price, portionSize float64) *error_wrapper.ErrorWrapper {
	err := i.itemDomain.Create(name, supplierID, category, unit, price, portionSize)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	items, err := i.itemDomain.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemService) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	item, err := i.itemDomain.FindByID(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemService) Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	err := i.itemDomain.Update(id, supplierID, name, category, price, unit)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := i.itemDomain.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
