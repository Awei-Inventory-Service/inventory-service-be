package item

import (
	"fmt"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (i *itemRepository) Create(name, supplierId string, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	item := model.Item{
		Name:       name,
		Category:   category,
		Price:      price,
		Unit:       unit,
		SupplierID: supplierId,
	}

	result := i.db.Create(&item)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (i *itemRepository) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	var items []model.Item
	result := i.db.Preload("Supplier").Find(&items)
	if result.Error != nil {
		fmt.Println("INI ERROR", result.Error)
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return items, nil
}

func (i *itemRepository) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	var item model.Item
	result := i.db.Where("uuid = ?", id).First(&item)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &item, nil
}

func (i *itemRepository) Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper {
	item := model.Item{
		Name:       name,
		Category:   category,
		Price:      price,
		Unit:       unit,
		SupplierID: supplierID,
	}

	result := i.db.Where("uuid = ?", id).Updates(&item)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (i *itemRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := i.db.Where("uuid = ?", id).Delete(&model.Item{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
