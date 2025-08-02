package sales

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type SalesResource interface {
	Create(sale model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper)
	Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type salesResource struct {
	db *gorm.DB
}
