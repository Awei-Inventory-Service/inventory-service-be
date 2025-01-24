package supplier

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(name, phoneNumber, address, picName string) error
	FindAll() ([]model.Supplier, error)
	FindByID(id string) (*model.Supplier, error)
	Update(id, name, phoneNumber, address, picName string) error
	Delete(id string) error
}

type supplierRepository struct {
	db *gorm.DB
}
