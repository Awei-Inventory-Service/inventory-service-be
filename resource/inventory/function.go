package inventory

import (
	"context"
	"errors"
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

func (i *inventoryResource) Create(inventory model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	result := i.db.Create(&inventory)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return &inventory, nil
}

func (i *inventoryResource) FindAll() ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	var inventories []model.Inventory
	result := i.db.Preload("Branch").Preload("Item").Find(&inventories)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return inventories, nil
}

func (i *inventoryResource) FindByBranch(branchID string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	var inventories []model.Inventory
	result := i.db.Where("branch_id = ?", branchID).Find(&inventories)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return inventories, nil
}

func (i *inventoryResource) FindByItem(itemID string) ([]model.Inventory, *error_wrapper.ErrorWrapper) {
	var inventory []model.Inventory
	result := i.db.Where("item_id = ?", itemID).Find(&inventory)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return inventory, nil
}

func (i *inventoryResource) FindByBranchAndItem(branchID, itemID string) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	var inventory model.Inventory

	result := i.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&inventory)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Stock balance record not found")
		}
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &inventory, nil
}

func (i *inventoryResource) Update(ctx context.Context, payload model.Inventory) (*model.Inventory, *error_wrapper.ErrorWrapper) {
	var result *gorm.DB

	if payload.UUID != "" {
		fmt.Println("Updating branch item based on UUID")
		result = i.db.WithContext(ctx).Where("uuid = ?", payload.UUID).Select("stock", "value").Updates(&payload)
	} else {
		fmt.Println("Updating branch item based on branch id and item id")
		fmt.Println("PAYLOAD", payload)
		result = i.db.WithContext(ctx).Where("branch_id = ? AND item_id = ?", payload.BranchID, payload.ItemID).Select("stock", "value").Updates(&payload)
	}

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return &payload, nil
}

func (i *inventoryResource) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	result := i.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.Inventory{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
