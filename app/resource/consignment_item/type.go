package consignmentitem

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type ConsignmentItemResource interface {
	Create(item model.ConsignmentItem) *error_wrapper.ErrorWrapper
	FindAll() ([]model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.ConsignmentItem, *error_wrapper.ErrorWrapper)
	Update(id string, item model.ConsignmentItem) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type consignmentItemResource struct {
	db *gorm.DB
}
