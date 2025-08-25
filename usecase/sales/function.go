package sales

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesService) Create(ctx context.Context, payload dto.CreateSalesRequest) *error_wrapper.ErrorWrapper {
	var (
		sales model.Sales
	)
	product, errW := s.productDomain.FindByID(ctx, payload.ProductID)

	if errW != nil {
		fmt.Println(errW.StackTrace(), errW.ActualError())
		return errW
	}
	var (
		updatedItemPurchaseChainDocuments []model.ItemPurchaseChainGet
		itemPurchaseWithSales             []model.ItemPurchaseChainGet
		totalCost                         float64
	)
	for _, ingredient := range product.Ingredients {
		cost, itemPurchaseChainDocuments, errW := s.itemPurchaseChainService.CalculateCost(
			ctx,
			ingredient.ItemID,
			payload.BranchID,
			// ingredient.Quantity*payload.Quantity,
			0,
		)

		if errW != nil {
			fmt.Println(errW.StackTrace(), errW.ActualError())
			return errW
		}

		updatedItemPurchaseChainDocuments = append(updatedItemPurchaseChainDocuments, itemPurchaseChainDocuments...)
		totalCost += cost
	}

	sales.Cost = totalCost
	sales.BranchID = payload.BranchID
	sales.ProductID = payload.ProductID
	sales.Quantity = payload.Quantity
	sales.Type = payload.Type

	newSales, errW := s.salesDomain.Create(sales)

	if errW != nil {
		fmt.Println(errW.StackTrace(), errW.ActualError())
		return errW
	}

	for _, doc := range updatedItemPurchaseChainDocuments {
		doc.Sales = append(doc.Sales, newSales.UUID)
	}

	// TO DO : Retry mechanism maybe
	errW = s.itemPurchaseChainDomain.BulkUpdate(ctx, itemPurchaseWithSales)

	if errW != nil {
		fmt.Println(errW.StackTrace(), errW.ActualError())
		return errW
	}

	return nil
}
