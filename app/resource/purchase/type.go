package purchase

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type PurchaseResource interface {
	Create(supplierId string, purchase model.Purchase) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id string, purchase model.Purchase) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseResource struct {
	db *gorm.DB
}
