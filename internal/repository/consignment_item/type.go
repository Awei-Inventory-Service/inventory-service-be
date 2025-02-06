package consignmentitem

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type ConsignmentItemRepository interface {
	Create(branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	Update(id, branchID, name string, quantity int, cost, price float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type consignmentItemRepository struct {
	db *gorm.DB
}
