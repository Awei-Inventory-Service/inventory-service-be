package supplier

import (
	"github.com/inventory-service/domain/supplier"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type SupplierService interface {
	Create(name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper
	FindByID(id string) (*model.Supplier, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Supplier, *error_wrapper.ErrorWrapper)
	Update(id, name, phoneNumber, address, picName string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type supplierService struct {
	supplierDomain supplier.SupplierDomain
}
