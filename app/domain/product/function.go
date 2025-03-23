package product

import (
	"context"

	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (p *productDomain) Create(ctx context.Context, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper {
	product := model.Product{
		Name:        name,
		Ingredients: ingredients,
	}

	return p.productResource.Create(ctx, product)
}

func (p *productDomain) FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper) {
	return p.productResource.FindAll(ctx)
}

func (p *productDomain) Find(ctx context.Context, payload model.GetProduct) ([]model.GetProduct, *error_wrapper.ErrorWrapper) {
	return p.productResource.Find(ctx, payload)
}

func (p *productDomain) FindByID(ctx context.Context, productID string) (model.Product, *error_wrapper.ErrorWrapper) {
	return p.productResource.FindByID(ctx, productID)
}

func (p *productDomain) Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper {
	return p.productResource.Update(ctx, productID, name, ingredients)
}

func (p *productDomain) Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper {
	return p.productResource.Delete(ctx, productID)
}
