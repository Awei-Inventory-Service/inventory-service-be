package consignmentitem

import (
	"github.com/inventory-service/internal/model"
)

func (c *consignmentItemRepository) Create(branchID, name string, quantity int, cost, price float64) error {
	item := model.ConsignmentItem{
		BranchID: branchID,
		Name:     name,
		Quantity: quantity,
		Cost:     cost,
		Price:    price,
	}

	result := c.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *consignmentItemRepository) FindAll() ([]model.ConsignmentItem, error) {
	var items []model.ConsignmentItem
	result := c.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (c *consignmentItemRepository) FindByID(id string) (*model.ConsignmentItem, error) {
	var item model.ConsignmentItem
	result := c.db.Where("uuid = ?", id).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (c *consignmentItemRepository) Update(id, branchID, name string, quantity int, cost, price float64) error {
	item := model.ConsignmentItem{
		BranchID: branchID,
		Name:     name,
		Quantity: quantity,
		Cost:     cost,
		Price:    price,
	}

	result := c.db.Where("uuid = ?", id).Updates(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *consignmentItemRepository) Delete(id string) error {
	result := c.db.Where("uuid = ?", id).Delete(&model.ConsignmentItem{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
