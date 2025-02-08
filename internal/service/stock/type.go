package stock

import (
	stockTransactionRepository "github.com/inventory-service/internal/repository/stock_transaction"
	"github.com/inventory-service/internal/service/model"
	"github.com/inventory-service/lib/error_wrapper"
)

type StockService interface {
	GetStockByItemID(itemID string) (model.Stock, *error_wrapper.ErrorWrapper)
}

type stockService struct {
	stockTransactionRepository stockTransactionRepository.StockTransactionRepository
}
