package item

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/repository/item"
)

type ItemService interface {
	Create(name, category string, price float64, unit string) error
	FindAll() ([]model.Item, error)
	FindByID(id string) (*model.Item, error)
	Update(id, name, category string, price float64, unit string) error
	Delete(id string) error
}

type itemService struct {
	itemRepository item.ItemRepository
}
