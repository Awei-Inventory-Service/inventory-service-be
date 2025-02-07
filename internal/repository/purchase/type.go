package purchase

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type PurchaseRepository interface {
	Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseRepository struct {
	db *gorm.DB
}
