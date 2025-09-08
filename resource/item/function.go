package item

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (i *itemResource) Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {

	result := i.db.Create(&item)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &item, nil
}

func (i *itemResource) FindAll() ([]model.Item, *error_wrapper.ErrorWrapper) {
	var items []model.Item
	result := i.db.
		Preload("ChildCompositions").
		Preload("ChildCompositions.ChildItem").
		Find(&items)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return items, nil
}

func (i *itemResource) FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper) {
	var item model.Item
	result := i.db.Where("uuid = ?", id).
		Preload("ChildCompositions").
		Preload("ChildCompositions.ChildItem").
		First(&item)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &item, nil
}

func (i *itemResource) Update(ctx context.Context, item model.Item) (*model.Item, *error_wrapper.ErrorWrapper) {
	if item.SupplierID != nil && *item.SupplierID == "" {
		item.SupplierID = nil
	}

	result := i.db.WithContext(ctx).Where("uuid = ?", item.UUID).Updates(&item)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return &item, nil
}

func (i *itemResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := i.db.Where("uuid = ?", id).Delete(&model.Item{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
