package purchase

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type PurchaseResource interface {
	Create(supplierId string, purchase model.Purchase) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	FindByItemID(itemID string) ([]model.Purchase, *error_wrapper.ErrorWrapper)
	FindByBranchAndItem(branchID, itemID string, offset, limit int) ([]model.Purchase, *error_wrapper.ErrorWrapper)
	Update(id string, purchase model.Purchase) *error_wrapper.ErrorWrapper
	Delete(id string) (*model.Purchase, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) (results []model.Purchase, errW *error_wrapper.ErrorWrapper)
}

type purchaseResource struct {
	db *gorm.DB
}
