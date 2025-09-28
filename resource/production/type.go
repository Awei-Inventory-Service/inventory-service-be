package production

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductionResource interface {
	Create(ctx context.Context, production model.Production) (*model.Production, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Production, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, payload model.Production) ([]model.Production, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Production, *error_wrapper.ErrorWrapper)
	Update(id string, production model.Production) *error_wrapper.ErrorWrapper
	Delete(id string) *error_wrapper.ErrorWrapper
}

type productionResource struct {
	db *gorm.DB
}
