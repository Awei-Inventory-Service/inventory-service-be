package stockbalance

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type StockBalanceRepository interface {
	Create(branchID, itemID string, currentStock int) error
	FindAll() ([]model.StockBalance, error)
	FindByBranch(branchID string) ([]model.StockBalance, error)
	FindByItem(itemID string) ([]model.StockBalance, error)
	FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, error)
	Update(branchID, itemID string, currentStock int) error
	Delete(branchID, itemID string) error
}

type stockBalanceRepository struct {
	db *gorm.DB
}
