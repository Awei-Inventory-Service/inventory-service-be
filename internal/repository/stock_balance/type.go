package stockbalance

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type StockBalanceRepository interface {
	Create(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper
	FindAll() ([]model.StockBalance, *error_wrapper.ErrorWrapper)
	FindByBranch(branchID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper)
	FindByItem(itemID string) ([]model.StockBalance, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string) (*model.StockBalance, *error_wrapper.ErrorWrapper)
	Update(branchID, itemID string, currentStock int) *error_wrapper.ErrorWrapper
	Delete(branchID, itemID string) *error_wrapper.ErrorWrapper
}

type stockBalanceRepository struct {
	db *gorm.DB
}
