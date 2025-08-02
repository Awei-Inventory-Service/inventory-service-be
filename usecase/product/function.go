package product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productService) Create(ctx context.Context, name string, ingredientsDto []dto.Ingredient) *error_wrapper.ErrorWrapper {
	var ingredients []model.Ingredient

	for _, ingredient := range ingredientsDto {
		item, err := p.itemDomain.FindByID(ingredient.ItemID)
		if err != nil {
			return error_wrapper.New(model.RErrMongoDBReadDocument, err.Error())
		}
		ingredients = append(ingredients, model.Ingredient{
			ItemID:   ingredient.ItemID,
			ItemName: item.Name,
			Quantity: ingredient.Quantity,
			Unit:     ingredient.Unit,
		})
	}
	err := p.productDomain.Create(ctx, name, ingredients)

	if err != nil {
		return err
	}
	return nil
}

func (p *productService) FindAll(ctx context.Context) ([]model.Product, *error_wrapper.ErrorWrapper) {
	products, err := p.productDomain.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) FindByID(ctx context.Context, productID string) (model.Product, *error_wrapper.ErrorWrapper) {
	product, err := p.productDomain.FindByID(ctx, productID)

	if err != nil {
		return model.Product{}, err
	}

	return product, err
}

func (p *productService) Update(ctx context.Context, productID string, name string, ingredients []model.Ingredient) *error_wrapper.ErrorWrapper {
	return p.productDomain.Update(ctx, productID, name, ingredients)
}

func (p *productService) Delete(ctx context.Context, producID string) *error_wrapper.ErrorWrapper {
	return p.productDomain.Delete(ctx, producID)
}
