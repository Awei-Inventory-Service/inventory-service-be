package stock

import (
	stockTransactionDomain "github.com/inventory-service/app/domain/stock_transaction"
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

type StockService interface {
	GetStockByItemID(itemID string) (model.Stock, *error_wrapper.ErrorWrapper)
}

type stockService struct {
	stockTransactionDomain stockTransactionDomain.StockTransactionDomain
}
