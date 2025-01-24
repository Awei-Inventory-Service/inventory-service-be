package stocktransaction

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type StockTransactionRepository interface {
	Create(transaction model.StockTransaction) error
	FindAll() ([]model.StockTransaction, error)
	FindByID(id string) (*model.StockTransaction, error)
	Update(id string, transaction model.StockTransaction) error
	Delete(id string) error
}

type stockTransactionRepository struct {
	db *gorm.DB
}
