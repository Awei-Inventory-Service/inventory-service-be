package item_branch

import (
	"errors"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

func (s *itemBranchResource) Create(itemBranch model.ItemBranch) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&itemBranch)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *itemBranchResource) FindAll() ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	var balances []model.ItemBranch
	result := s.db.Preload("Branch").Preload("Item").Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *itemBranchResource) FindByBranch(branchID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	var balances []model.ItemBranch
	result := s.db.Where("branch_id = ?", branchID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *itemBranchResource) FindByItem(itemID string) ([]model.ItemBranch, *error_wrapper.ErrorWrapper) {
	var balances []model.ItemBranch
	result := s.db.Where("item_id = ?", itemID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *itemBranchResource) FindByBranchAndItem(branchID, itemID string) (*model.ItemBranch, *error_wrapper.ErrorWrapper) {
	var balance model.ItemBranch

	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&balance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Stock balance record not found")
		}
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &balance, nil
}

func (s *itemBranchResource) Update(branchID, itemID string, currentStock float64) *error_wrapper.ErrorWrapper {
	result := s.db.Model(&model.ItemBranch{}).
		Where("branch_id = ? AND item_id = ?", branchID, itemID).
		Update("current_stock", currentStock)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *itemBranchResource) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.ItemBranch{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
