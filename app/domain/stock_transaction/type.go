package stocktransaction

import (
	"github.com/inventory-service/app/model"
	stocktransaction "github.com/inventory-service/app/resource/stock_transaction"
	"github.com/inventory-service/lib/error_wrapper"
)

type StockTransactionDomain interface {
	Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper)
	Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
	FindWithFilter(filters []map[string]interface{}) ([]model.StockTransaction, *error_wrapper.ErrorWrapper)
}

type stockTransactionDomain struct {
	stockTransactionResource stocktransaction.StockTransactionResource
}
