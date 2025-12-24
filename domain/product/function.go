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
		product.Code = rawProduct.Code

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

		for _, branch := range rawProduct.BranchProducts {
			product.Branches = append(product.Branches, dto.GetProductBranchResponse{
				BranchID:           branch.BranchID,
				BranchName:         branch.Branch.Name,
				BranchProductPrice: *branch.SellingPrice,
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

func (p *productDomain) CalculateProductCost(
	ctx context.Context,
	product model.Product,
	branchID string,
	timestamp time.Time,
) (
	results []dto.ProductRecipeWithPrice,
	totalCost float64,
	errW *error_wrapper.ErrorWrapper,
) {

	for _, productComposition := range product.ProductRecipe {
		item, errW := p.itemResource.FindByID(productComposition.ItemID)

		if errW != nil {
			fmt.Println("Error getting item", errW)
			return results, totalCost, errW
		}

		_, _, errW = p.inventoryDomain.SyncBranchItem(ctx, branchID, productComposition.ItemID)

		if errW != nil {
			return results, totalCost, errW
		}
		productCompositionAmount := utils.StandarizeMeasurement(productComposition.Amount, productComposition.Unit, item.Unit)

		inventorySnapshot, errW := p.inventorySnapshotResource.Get(ctx, []dto.Filter{
			{
				Key:    "item_id",
				Values: []string{productComposition.ItemID},
			},
			{
				Key:    "day",
				Values: []string{fmt.Sprintf("%d", timestamp.Day())},
			},
			{
				Key:    "month",
				Values: []string{fmt.Sprintf("%d", int(timestamp.Month()))},
			},
			{
				Key:    "year",
				Values: []string{fmt.Sprintf("%d", timestamp.Year())},
			},
		},
			[]dto.Order{}, 1, 0)

		if errW != nil {
			if errW.Is(model.RErrDataNotFound) {
				_, balance, price, errW := p.inventoryDomain.CalculatePriceAndBalance(ctx, timestamp, productComposition.ItemID, branchID, nil)
				if errW != nil {
					fmt.Println("Error calculating price and balance in CalculateProductCost", errW)
					continue
				}

				errW = p.inventorySnapshotResource.Upsert(ctx, dto.CreateInventorySnapshotRequest{
					Balance:  balance,
					Value:    price,
					BranchID: branchID,
					ItemID:   productComposition.ItemID,
					Date:     timestamp,
				})
				if errW != nil {
					fmt.Println("Error upserting inventory snapshot resource ", errW)
					continue
				}
				results = append(results, dto.ProductRecipeWithPrice{
					UUID:   productComposition.UUID,
					ItemID: productComposition.ItemID,
					Cost:   price * productCompositionAmount,
					Unit:   productComposition.Unit,
					Amount: productComposition.Amount,
				})
				continue
			}
			fmt.Println("Error getting inventory snapshot")
			return results, totalCost, errW
		}

		itemValue := inventorySnapshot[0].Latest
		fmt.Println("Product composition amount", productCompositionAmount, itemValue)

		results = append(results, dto.ProductRecipeWithPrice{
			UUID:   productComposition.UUID,
			ItemID: productComposition.ItemID,
			Cost:   itemValue * productCompositionAmount,
			Unit:   productComposition.Unit,
			Amount: productComposition.Amount,
		})
		totalCost += (itemValue * productCompositionAmount)
	}

	errW = p.productSnapshotResource.Upsert(ctx, dto.CreateProductSnapshotRequest{
		ProductID: product.UUID,
		BranchID:  branchID,
		Value:     totalCost,
	})

	if errW != nil {
		fmt.Println("Error upserting product snapshot resource", errW)
		return
	}
	return
}
