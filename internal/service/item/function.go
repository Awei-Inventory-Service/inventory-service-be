package item

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemService) Create(name, supplierID string, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	err := i.itemRepository.Create(name, supplierID, category, price, unit)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	items, err := i.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemService) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	item, err := i.itemRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemService) Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	err := i.itemRepository.Update(id, supplierID, name, category, price, unit)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) Delete(id string) *error_wrapper.ErrorWrapper {
	err := i.itemRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
