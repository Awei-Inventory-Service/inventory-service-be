package product

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productService) Create(ctx context.Context, payload dto.CreateProductRequest) *error_wrapper.ErrorWrapper {

	productType := payload.MapProductCategory()

	if productType == "" {
		return error_wrapper.New(model.UErrInvalidProductType, "Invalid product type")
	}

	product, errW := p.productDomain.Create(ctx, model.Product{
		Name:         payload.Name,
		Code:         payload.Code,
		Category:     payload.Category,
		Type:         productType,
		Unit:         payload.Unit,
		SellingPrice: payload.SellingPrice,
	})

	if errW != nil {
		return errW
	}

	for _, composition := range payload.ProductRecipes {
		errW = p.productRecipeDomain.Create(ctx, model.ProductRecipe{
			ProductID: product.UUID,
			ItemID:    composition.ItemID,
			Amount:    composition.Amount,
			Unit:      composition.Unit,
			Notes:     composition.Notes,
		})

		if errW != nil {
			return errW
		}
	}

	for _, branchId := range payload.BranchIDs {
		_, errW = p.branchProductDomain.Create(ctx, dto.CreateBranchProductRequest{
			BranchID:     branchId,
			ProductID:    product.UUID,
			SellingPrice: payload.SellingPrice,
		})

		if errW != nil {
			return errW
		}
	}

	return nil
}

func (p *productService) FindAll(ctx context.Context) ([]dto.GetProductResponse, *error_wrapper.ErrorWrapper) {
	products, err := p.productDomain.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productService) FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper) {
	product, err := p.productDomain.FindByID(ctx, productID)

	if err != nil {
		return nil, err
	}

	return product, err
}

func (p *productService) Update(ctx context.Context, product dto.UpdateProductRequest, productID string) *error_wrapper.ErrorWrapper {

	return p.productDomain.Update(ctx, product, productID)
}

func (p *productService) Delete(ctx context.Context, producID string) *error_wrapper.ErrorWrapper {
	return p.productDomain.Delete(ctx, producID)
}

func (p *productService) GetProductCost(ctx context.Context, productID, branchID string) (cost float64, errW *error_wrapper.ErrorWrapper) {
	product, errW := p.productDomain.FindByID(ctx, productID)

	if errW != nil {
		return
	}

	_, cost, errW = p.productDomain.CalculateProductCost(ctx, product.ProductRecipe, branchID)

	return
}
