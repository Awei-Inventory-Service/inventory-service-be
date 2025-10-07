package product

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/utils"
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
		product.Unit = rawProduct.Unit
		product.Category = rawProduct.Category
		product.SellingPrice = rawProduct.SellingPrice

		for _, ingredient := range rawProduct.ProductRecipe {
			item, errW := p.itemResource.FindByID(ingredient.ItemID)

			if errW != nil {
				return nil, errW
			}

			product.Ingredients = append(product.Ingredients, dto.GetIngredient{
				ItemID:      ingredient.ItemID,
				ItemName:    item.Name,
				ItemUnit:    ingredient.Unit,
				ItemPortion: ingredient.Amount,
			})

		}
		productsResponse = append(productsResponse, product)

	}
	return productsResponse, nil
}

func (p *productDomain) FindByID(ctx context.Context, productID string) (*model.Product, *error_wrapper.ErrorWrapper) {
	return p.productResource.FindByID(ctx, productID)
}

func (p *productDomain) Update(ctx context.Context, payload dto.UpdateProductRequest, productID string) *error_wrapper.ErrorWrapper {
	product := model.Product{
		UUID:         productID,
		Code:         payload.Code,
		Name:         payload.Name,
		Category:     payload.Category,
		Unit:         payload.Unit,
		SellingPrice: payload.SellingPrice,
		UpdatedAt:    time.Now(),
	}

	updatedProduct, errW := p.productResource.Update(ctx, product)

	if errW != nil {
		return errW
	}

	errW = p.productRecipeResource.DeleteByProductID(ctx, updatedProduct.UUID)

	if errW != nil {
		return errW
	}

	for _, productComposition := range payload.ProductRecipes {
		errW = p.productRecipeResource.Create(ctx, model.ProductRecipe{
			ProductID: updatedProduct.UUID,
			Amount:    productComposition.Amount,
			Unit:      productComposition.Unit,
			Notes:     productComposition.Notes,
			ItemID:    productComposition.ItemID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if errW != nil {
			return errW
		}
	}

	return nil
}

func (p *productDomain) Delete(ctx context.Context, productID string) *error_wrapper.ErrorWrapper {
	return p.productResource.Delete(ctx, productID)
}

func (p *productDomain) CalculateProductCost(ctx context.Context, productCompositions []model.ProductRecipe, branchID string) (float64, *error_wrapper.ErrorWrapper) {
	var (
		price float64
	)

	for _, productComposition := range productCompositions {
		_, _, errW := p.inventoryDomain.SyncBranchItem(ctx, branchID, productComposition.ItemID)

		if errW != nil {
			return 0.0, errW
		}

		inventory, errW := p.inventoryResource.FindByBranchAndItem(branchID, productComposition.ItemID)
		if errW != nil {
			fmt.Printf("Error finding branch item with branch_idx: %s and item_id: %s ", branchID, productComposition.ItemID)
			return price, errW
		}
		fmt.Println("Product composition amount", productComposition.Amount, productComposition.Unit, inventory.Item.Unit)

		productCompositionAmount := utils.StandarizeMeasurement(productComposition.Amount, productComposition.Unit, inventory.Item.Unit)
		fmt.Println("Product composition amount", productCompositionAmount, inventory.Value)
		price += (inventory.Value * productCompositionAmount)
	}

	return price, nil
}
