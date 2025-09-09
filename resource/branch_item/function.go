package branch_item

import (
	"context"
	"errors"
	"fmt"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

func (s *branchItemResource) Create(branchItem model.BranchItem) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&branchItem)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *branchItemResource) FindAll() ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	var branchItems []model.BranchItem
	result := s.db.Preload("Branch").Preload("Item").Find(&branchItems)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return branchItems, nil
}

func (s *branchItemResource) FindByBranch(branchID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	var branchItems []model.BranchItem
	result := s.db.Where("branch_id = ?", branchID).Find(&branchItems)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return branchItems, nil
}

func (s *branchItemResource) FindByItem(itemID string) ([]model.BranchItem, *error_wrapper.ErrorWrapper) {
	var branchItem []model.BranchItem
	result := s.db.Where("item_id = ?", itemID).Find(&branchItem)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return branchItem, nil
}

func (s *branchItemResource) FindByBranchAndItem(branchID, itemID string) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	var branchItem model.BranchItem

	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&branchItem)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Stock balance record not found")
		}
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &branchItem, nil
}

func (s *branchItemResource) Update(ctx context.Context, payload model.BranchItem) (*model.BranchItem, *error_wrapper.ErrorWrapper) {
	var result *gorm.DB

	if payload.UUID != "" {
		fmt.Println("Updating branch item based on UUID")
		result = s.db.WithContext(ctx).Where("uuid = ?", payload.UUID).Select("current_stock", "price").Updates(&payload)
	} else {
		fmt.Println("Updating branch item based on branch id and item id")
		fmt.Println("PAYLOAD", payload)
		result = s.db.WithContext(ctx).Where("branch_id = ? AND item_id = ?", payload.BranchID, payload.ItemID).Select("current_stock", "price").Updates(&payload)
	}

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return &payload, nil
}

func (s *branchItemResource) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.BranchItem{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
