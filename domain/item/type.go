package item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/item"
)

type ItemDomain interface {
	Create(name, supplierID, unit string, category model.ItemCategory, price, portionSize float64) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper)
	Update(id, supplierID, name, unit string, category model.ItemCategory, price float64) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemDomain struct {
	itemResource item.ItemResource
}
