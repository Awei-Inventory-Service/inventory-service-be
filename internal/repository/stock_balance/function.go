package stockbalance

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *stockBalanceRepository) Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	stockBalance := model.StockBalance{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: currentStock,
	}

	result := s.db.Create(&stockBalance)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockBalanceRepository) FindAll() ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByBranch(branchID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Where("branch_id = ?", branchID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByItem(itemID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balances []model.StockBalance
	result := s.db.Where("item_id = ?", itemID).Find(&balances)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, *error_wrapper.ErrorWrapper) {
	var balance model.StockBalance
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&balance)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &balance, nil
}

func (s *stockBalanceRepository) Update(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper {
	result := s.db.Model(&model.StockBalance{}).
		Where("branch_id = ? AND item_id = ?", branchID, itemID).
		Update("current_stock", currentStock)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockBalanceRepository) Delete(branchID, itemID string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.StockBalance{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
