package productionitem

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"gorm.io/gorm"
)

type ProductionItemResource interface {
	Create(ctx context.Context, productionItem model.ProductionItem) (*model.ProductionItem, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter model.ProductionItem) ([]model.ProductionItem, *error_wrapper.ErrorWrapper)
}

type productionItemResource struct {
	db *gorm.DB
}
