package production

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductionResource interface {
	Create(ctx context.Context, production model.Production) (*model.Production, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Production, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]model.Production, int64, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Production, *error_wrapper.ErrorWrapper)
	Update(id string, production model.Production) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, id string) *error_wrapper.ErrorWrapper
}

type productionResource struct {
	db *gorm.DB
}
