package product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productDomain) Create(ctx context.Context, payload model.Product) (*model.Product, *error_wrapper.ErrorWrapper) {

	return p.productResource.Create(ctx, payload)
}

func (p *productDomain) FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper) {
	var productsResponse []dto.GetProductResponse
	products, errW := p.productResource.FindAll(ctx)

	if errW != nil {
		return nil, errW
	}

	for _, rawProduct := range products {
		var product dto.GetProductResponse
		product.Id = rawProduct.UUID
		product.Name = rawProduct.Name
		for _, ingredient := range rawProduct.ProductComposition {
			item, errW := p.itemResource.FindByID(ingredient.ItemID)

			if errW != nil {
				return nil, errW
			}

			product.Ingredients = append(product.Ingredients, dto.GetIngredient{
				ItemID:      ingredient.ItemID,
				ItemName:    item.Name,
				ItemPortion: item.PortionSize,
				Ratio:       ingredient.Ratio,
				ItemUnit:    item.Unit,
			})

			productsResponse = append(productsResponse, product)
		}
	}
	return productsResponse, nil
}

func (p *productDomain) FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper) {
	return p.productResource.FindByID(ctx, productID)
}

func (p *productDomain) Update(ctx context.Context, payload model.Product) *error_wrapper.ErrorWrapper {
	return p.productResource.Update(ctx, payload)
}

func (p *productDomain) Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper {
	return p.productResource.Delete(ctx, productID)
}
