package stocktransaction

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *stockTransactionResource) Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Create(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transactions []model.StockTransaction
	result := s.db.Find(&transactions)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return transactions, nil
}

func (s *stockTransactionResource) FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transaction model.StockTransaction
	result := s.db.Where("uuid = ?", id).First(&transaction)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &transaction, nil
}

func (s *stockTransactionResource) Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Updates(&transaction)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) Delete(id string) *error_wrapper.ErrorWrapper {
	result := s.db.Where("uuid = ?", id).Delete(&model.StockTransaction{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}

	return nil
}

func (s *stockTransactionResource) FindWithFilter(filters []map[string]interface{}, sort string) ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	var transactions []model.StockTransaction
	query := s.db

	for _, filter := range filters {
		field, okField := filter["field"].(string)
		value, okValue := filter["value"]
		if okField && okValue {
			query = query.Where(field+" = ?", value)
		}
	}

	if sort != "" {
		query = query.Order(sort)
	}

	result := query.Preload("Item").Find(&transactions)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return transactions, nil
}
