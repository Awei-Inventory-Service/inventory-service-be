package stocktransaction

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *stockTransactionRepository) Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionRepository) FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transactions []model.StockTransaction
	result := s.db.Find(&transactions)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return transactions, nil
}

func (s *stockTransactionRepository) FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transaction model.StockTransaction
	result := s.db.Where("uuid = ?", id).First(&transaction)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &transaction, nil
}

func (s *stockTransactionRepository) Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionRepository) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.StockTransaction{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}
