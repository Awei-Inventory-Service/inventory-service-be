package purchase

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/purchase"
)

type PurchaseDomain interface {
	Create(supplierId, branchId, itemId string, quantity int, purchaseCost float64) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id, supplierId, branchId, itemId string, quantity int, purchaseCost float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type purchaseDomain struct {
	purchaseResource purchase.PurchaseResource
}
