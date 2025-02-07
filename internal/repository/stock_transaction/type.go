package stocktransaction

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type StockTransactionRepository interface {
	Create(transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	FindAll() ([]model.StockTransaction, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.StockTransaction, *error_wrapper.ErrorWrapper)
	Update(id string, transaction model.StockTransaction) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type stockTransactionRepository struct {
	db *gorm.DB
}
