package item

import (
	"github.com/inventory-service/internal/model"
)

func (i *itemService) Create(name, category string, price float64, unit string) error {
	err := i.itemRepository.Create(name, category, price, unit)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) FindAll() ([]model.Item, error) {
	items, err := i.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (i *itemService) FindByID(id string) (*model.Item, error) {
	item, err := i.itemRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *itemService) Update(id, name, category string, price float64, unit string) error {
	err := i.itemRepository.Update(id, name, category, price, unit)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) Delete(id string) error {
	err := i.itemRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
