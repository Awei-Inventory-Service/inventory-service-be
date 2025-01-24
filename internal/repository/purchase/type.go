package purchase

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type PurchaseRepository interface {
	Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) error
	FindAll() ([]model.Purchase, error)
	FindByID(id string) (*model.Purchase, error)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) error
	Delete(id string) error
}

type purchaseRepository struct {
	db *gorm.DB
}
