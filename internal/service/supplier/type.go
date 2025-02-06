package supplier

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/supplier"
	"github.com/inventory-service/lib/error_wrapper"
)

type SupplierService interface {
	Create(name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper
	FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper)
	Update(id, name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type supplierService struct {
	supplierRepository supplier.SupplierRepository
}
