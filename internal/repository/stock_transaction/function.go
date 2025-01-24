package stocktransaction

import (
	"github.com/inventory-service/internal/model"
)

func (s *stockTransactionRepository) Create(transaction model.StockTransaction) error {
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *stockTransactionRepository) FindAll() ([]model.StockTransaction, error) {
	var transactions []model.StockTransaction
	result := s.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (s *stockTransactionRepository) FindByID(id string) (*model.StockTransaction, error) {
	var transaction model.StockTransaction
	result := s.db.Where("uuid = ?", id).First(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	return &transaction, nil
}

func (s *stockTransactionRepository) Update(id string, transaction model.StockTransaction) error {
	result := s.db.Where("uuid = ?", id).Updates(&transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *stockTransactionRepository) Delete(id string) error {
	result := s.db.Where("uuid = ?", id).Delete(&model.StockTransaction{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
