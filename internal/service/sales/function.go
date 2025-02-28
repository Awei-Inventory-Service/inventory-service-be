package sales

import (
	"context"

	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (s *salesService) Create(ctx context.Context, payload model.Sales) *error_wrapper.ErrorWrapper {

	product, errW := s.productRepository.FindByID(ctx, payload.ProductID)

	if errW != nil {
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
			ingredient.Quantity*payload.Quantity,
		)

		if errW != nil {
			return errW
		}

		updatedItemPurchaseChainDocuments = append(updatedItemPurchaseChainDocuments, itemPurchaseChainDocuments...)
		totalCost += cost
	}

	payload.Cost = totalCost
	newSales, errW := s.salesRepository.Create(payload)

	if errW != nil {
		return errW
	}

	for _, doc := range updatedItemPurchaseChainDocuments {
		doc.Sales = append(doc.Sales, newSales.UUID)
	}

	// TO DO : Retry mechanism maybe
	errW = s.itemPurchaseChainRepository.BulkUpdate(ctx, itemPurchaseWithSales)

	if errW != nil {
		return errW
	}

	return nil
}
