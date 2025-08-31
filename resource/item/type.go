package item

import (
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ItemResource interface {
	Create(item model.Item) (*model.Item, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Item, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Item, *error_wrapper.ErrorWrapper)
	Update(id string, item model.Item) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type itemResource struct {
	db *gorm.DB
}
