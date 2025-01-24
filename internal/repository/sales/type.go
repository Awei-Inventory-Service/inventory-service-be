package sales

import (
	"github.com/inventory-service/internal/model"
	"gorm.io/gorm"
)

type SalesRepository interface {
	Create(sale model.Sales) error
	FindAll() ([]model.Sales, error)
	FindByID(id string) (*model.Sales, error)
	Update(id string, sale model.Sales) error
	Delete(id string) error
}

type salesRepository struct {
	db *gorm.DB
}
