package supplier

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/supplier"
)

type SupplierService interface {
	Create(name, phoneNumber, address, picName string) error
	FindByID(id string) (*model.Supplier, error)
	FindAll() ([]model.Supplier, error)
	Update(id, name, phoneNumber, address, picName string) error
	Delete(id string) error
}

type supplierService struct {
	supplierRepository supplier.SupplierRepository
}
