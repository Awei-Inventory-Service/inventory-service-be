package item

import (
	"github.com/inventory-service/internal/model"
)

func (i *itemRepository) Create(name, supplierId string, category string, price float64, unit string) error {
	item := model.Item{
		Name:       name,
		Category:   category,
		Price:      price,
		Unit:       unit,
		SupplierID: supplierId,
	}

	result := i.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (i *itemRepository) FindAll() ([]model.Item, error) {
	var items []model.Item
	result := i.db.Preload("Supplier").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (i *itemRepository) FindByID(id string) (*model.Item, error) {
	var item model.Item
	result := i.db.Where("uuid = ?", id).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (i *itemRepository) Update(id, supplierID string, name, category string, price float64, unit string) error {
	item := model.Item{
		Name:     name,
		Category: category,
		Price:    price,
		Unit:     unit,
		SupplierID: supplierID,
	}

	result := i.db.Where("uuid = ?", id).Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (i *itemRepository) Delete(id string) error {
	result := i.db.Where("uuid = ?", id).Delete(&model.Item{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
