package item

import (
	"github.com/inventory-service/domain/item"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

type ItemService interface {
	Create(name, supplierID, category, unit string, price, portionSize float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper)
	Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemService struct {
	itemDomain item.ItemDomain
}
