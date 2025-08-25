package stockbalance

import (
	"errors"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

func (s *stockBalanceResource) Create(stockBalance model.StockBalance) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&stockBalance)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockBalanceResource) FindAll() ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceResource) FindByBranch(branchID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Where("branch_id = ?", branchID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceResource) FindByItem(itemID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Where("item_id = ?", itemID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceResource) FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balance model.StockBalance
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&balance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, error_wrapper.New(model.RErrDataNotFound, "Stock balance record not found")
		}
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &balance, nil
}

func (s *stockBalanceResource) Update(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	result := s.db.Model(&model.StockBalance{}).
		Where("branch_id = ? AND item_id = ?", branchID, itemID).
		Update("current_stock", currentStock)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockBalanceResource) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.StockBalance{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
