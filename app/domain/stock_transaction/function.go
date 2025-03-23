package stocktransaction

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *stockTransactionDomain) Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Create(transaction)
}

func (s *stockTransactionDomain) FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindAll()
}

func (s *stockTransactionDomain) FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindByID(id)
}

func (s *stockTransactionDomain) Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Update(id, transaction)
}

func (s *stockTransactionDomain) Delete(id string) *error_wrapper.ErrorWrapper {
	return s.stockTransactionResource.Delete(id)
}

func (s *stockTransactionDomain) FindWithFilter(filters []map[string]interface{}) ([]model.StockTransaction, *error_wrapper.ErrorWrapper) {
	return s.stockTransactionResource.FindWithFilter(filters)
}
