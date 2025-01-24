package stockbalance

import "github.com/inventory-service/internal/model"

func (s *stockBalanceRepository) Create(branchID, itemID string, currentStock int) error {
	stockBalance := model.StockBalance{
		BranchID:     branchID,
		ItemID:       itemID,
		CurrentStock: currentStock,
	}

	result := s.db.Create(&stockBalance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *stockBalanceRepository) FindAll() ([]model.StockBalance, error) {
	var balances []model.StockBalance
	result := s.db.Find(&balances)
	if result.Error != nil {
		return nil, result.Error
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByBranch(branchID string) ([]model.StockBalance, error) {
	var balances []model.StockBalance
	result := s.db.Where("branch_id = ?", branchID).Find(&balances)
	if result.Error != nil {
		return nil, result.Error
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByItem(itemID string) ([]model.StockBalance, error) {
	var balances []model.StockBalance
	result := s.db.Where("item_id = ?", itemID).Find(&balances)
	if result.Error != nil {
		return nil, result.Error
	}

	return balances, nil
}

func (s *stockBalanceRepository) FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, error) {
	var balance model.StockBalance
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).First(&balance)
	if result.Error != nil {
		return nil, result.Error
	}

	return &balance, nil
}

func (s *stockBalanceRepository) Update(branchID, itemID string, currentStock int) error {
	result := s.db.Model(&model.StockBalance{}).
		Where("branch_id = ? AND item_id = ?", branchID, itemID).
		Update("current_stock", currentStock)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *stockBalanceRepository) Delete(branchID, itemID string) error {
	result := s.db.Where("branch_id = ? AND item_id = ?", branchID, itemID).Delete(&model.StockBalance{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
