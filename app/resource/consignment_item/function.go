package consignmentitem

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (c *consignmentItemResource) Create(item model.ConsignmentItem) *error_wrapper.ErrorWrapper {

	result := c.db.Create(&item)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (c *consignmentItemResource) FindAll() ([]model.ConsignmentItem, *error_wrapper.ErrorWrapper) {
	var items []model.ConsignmentItem
	result := c.db.Find(&items)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return items, nil
}

func (c *consignmentItemResource) FindByID(id string) (*model.ConsignmentItem, *error_wrapper.ErrorWrapper) {
	var item model.ConsignmentItem
	result := c.db.Where("uuid = ?", id).First(&item)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &item, nil
}

func (c *consignmentItemResource) Update(id string, item model.ConsignmentItem) *error_wrapper.ErrorWrapper {
	result := c.db.Where("uuid = ?", id).Updates(&item)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (c *consignmentItemResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := c.db.Where("uuid = ?", id).Delete(&model.ConsignmentItem{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
