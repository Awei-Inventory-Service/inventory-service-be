package supplier

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type SupplierResource interface {
	Create(supplier model.Supplier) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper)
	Update(id string, supplier model.Supplier) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type supplierResource struct {
	db *gorm.DB
}
