package sales

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesService) Create(ctx context.Context, payload dto.CreateSalesRequest, userID string) *error_wrapper.ErrorWrapper {
	var (
		sales         model.Sales
		branchProduct *model.BranchProduct
	)

	product, errW := s.productDomain.FindByID(ctx, payload.ProductID)

	if errW != nil {
		return errW
	}

	productCost, errW := s.productDomain.CalculateProductCost(ctx, product.ProductComposition, payload.BranchID)

	if errW != nil {
		return errW
	}

	sales.Cost = productCost

	branchProduct, errW = s.branchProductDomain.GetByBranchIdAndProductId(ctx, payload.BranchID, payload.ProductID)

	if errW != nil {
		if errW.Is(model.RErrDataNotFound) {
			// If not found, create 1
			errW = nil

			branchProduct, errW = s.branchProductDomain.Create(ctx, dto.CreateBranchProductRequest{
				BranchID:     payload.BranchID,
				ProductID:    payload.ProductID,
				SellingPrice: product.SellingPrice,
			})

			if errW != nil {
				return errW
			}
		} else {
			return errW
		}

	}

	sales.BranchProductID = branchProduct.UUID
	sales.Quantity = payload.Quantity
	sales.Type = payload.Type
	if branchProduct.SellingPrice != nil {
		sales.Price = *branchProduct.SellingPrice
	} else {
		sales.Price = 0.0
	}

	_, errW = s.salesDomain.Create(ctx, sales)

	if errW != nil {
		return errW
	}

	for _, itemComposition := range product.ProductComposition {
		total := itemComposition.Item.PortionSize * itemComposition.Ratio * payload.Quantity
		errW = s.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              itemComposition.ItemID,
			Type:                "OUT",
			IssuerID:            userID,
			Quantity:            total,
			// To-do: recheck this cost
			Cost: sales.Cost,
			Unit: itemComposition.Item.Unit,
		})
		
		if errW != nil {
			return errW
		}

	}
	return nil
}
