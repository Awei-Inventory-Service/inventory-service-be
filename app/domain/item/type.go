package item

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/item"
	"github.com/inventory-service/lib/error_wrapper"
)

type ItemDomain interface {
	Create(name, supplierID string, category string, price float64, unit string) *error_wrapper.ErrorWrapper
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper)
	Update(id, supplierID string, name, category string, price float64, unit string) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemDomain struct {
	itemResource item.ItemResource
}
