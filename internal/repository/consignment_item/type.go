package consignmentitem

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type ConsignmentItemRepository interface {
	Create(branchID, name string, quantity int, cost, price float64) error
	FindAll() ([]model.ConsignmentItem, error)
	FindByID(id string) (*model.ConsignmentItem, error)
	Update(id, branchID, name string, quantity int, cost, price float64) error
	Delete(id string) error
}

type consignmentItemRepository struct {
	db *gorm.DB
}
