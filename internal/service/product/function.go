package product

import (
	"context"

	"github.com/inventory-service/internal/model"
)

func (p *productService) Create(ctx context.Context, name string, ingredients []model.Ingredient) error {
	err := p.productRepository.Create(ctx, name, ingredients)

	if err != nil {
		return err
	}
	return nil
}

func (p *productService) FindAll(ctx context.Context) ([]model.Product, error) {
	products, err := p.productRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) FindByID(ctx context.Context, productID string) (model.Product, error) {
	product, err := p.productRepository.FindByID(ctx, productID)

	if err != nil {
		return model.Product{}, err
	}

	return product, err
}

func (p *productService) Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) error {
	return p.productRepository.Update(ctx, productID, name, ingredients)
}

func (p *productService) Delete(ctx context.Context, producID string) error {
	return p.productRepository.Delete(ctx, producID)
}
