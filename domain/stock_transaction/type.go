package stocktransaction

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	stocktransaction "github.com/inventory-service/resource/stock_transaction"
)

type StockTransactionDomain interface {
	Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper)
	Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
	FindWithFilter(filters []map[string]interface{}, sort string, limit, offset int) ([]model.StockTransaction, *error_wrapper.ErrorWrapper)
	InvalidateStockTransaction(ctx context.Context, filter []map[string]interface{}, userID string) (errW *error_wrapper.ErrorWrapper)
}

type stockTransactionDomain struct {
	stockTransactionResource stocktransaction.StockTransactionResource
}
