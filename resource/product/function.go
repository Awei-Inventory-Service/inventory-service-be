package product

import (
	"context"

	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productResource) Create(ctx context.Context, product model.Product) (*model.Product, *error_wrapper.ErrorWrapper) {

	result := p.db.Create(&product)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresCreateDocument, result.Error)
	}

	return &product, nil
}

func (p *productResource) FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper) {
	var products []model.Product

	result := p.db.WithContext(ctx).
		Preload("ProductComposition").
		Preload("ProductComposition.Item").
		Find(&products)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return products, nil
}

func (p *productResource) FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper) {
	var product model.Product

	result := p.db.Where("uuid = ?", productID).First(&product)

	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &product, nil
}

func (p *productResource) Update(ctx context.Context, product model.Product) (*model.Product, *error_wrapper.ErrorWrapper) {
	result := p.db.Where("uuid = ?", product.UUID).Updates(&product)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresUpdateDocument, result.Error.Error())
	}
	return &product, nil
}

func (p *productResource) Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper {

	result := p.db.Where("uuid = ? ", productID).Delete(&model.Product{})
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresDeleteDocument, result.Error.Error())
	}
	return nil
}
