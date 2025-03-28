package product

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/mongodb"
	"github.com/inventory-service/lib/error_wrapper"
)

type ProductResource interface {
	Create(ctx context.Context, product model.Product) *error_wrapper.ErrorWrapper
	FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper)
	FindByID(ctx context.Context, productID string) (model.Product, *error_wrapper.ErrorWrapper)
	Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper
	Find(ctx context.Context, payload model.GetProduct) ([]model.GetProduct, *error_wrapper.ErrorWrapper)
}

type productResource struct {
	productCollection mongodb.MongoDBCollectionWrapper
}
