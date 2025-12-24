package sales

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/constant"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesService) Create(ctx context.Context, payload dto.CreateSalesRequest, userID string) *error_wrapper.ErrorWrapper {
	var (
		branchProduct   *model.BranchProduct
		transactionDate time.Time
		sales           model.Sales
	)
	if payload.TransactionDate != "" {
		parsedDate, err := time.Parse("2006-01-02", payload.TransactionDate)
		if err != nil {
			return error_wrapper.New(model.CErrJsonBind, "Invalid date format. Expected YYYY-MM-DD")
		}
		transactionDate = parsedDate
	} else {
		transactionDate = time.Now()
	}

	sales.BranchID = payload.BranchID
	sales.TransactionDate = transactionDate
	newSales, errW := s.salesDomain.Create(ctx, sales)
	if errW != nil {
		fmt.Println("Error creating new sales ", errW)
		return errW
	}

	for _, sales := range payload.SalesData {
		var (
			salesProduct model.SalesProduct
		)

		product, errW := s.productDomain.FindByID(ctx, sales.ProductID)
		if errW != nil {
			return errW
		}

		productRecipes, totalPrice, errW := s.productDomain.CalculateProductCost(ctx, *product, payload.BranchID, transactionDate)
		if errW != nil {
			return errW
		}
		branchProduct, errW = s.branchProductDomain.GetByBranchIdAndProductId(ctx, payload.BranchID, sales.ProductID)
		if errW != nil {
			fmt.Println("Error getting branch product", errW)
			return errW
		}

		fmt.Printf("Product cost : %f, sales data quantity: %f", totalPrice, sales.Quantity)
		salesProduct.BranchID = payload.BranchID
		salesProduct.Cost = totalPrice * sales.Quantity
		salesProduct.ProductID = product.UUID
		salesProduct.Quantity = sales.Quantity
		salesProduct.Type = sales.Type
		salesProduct.SalesID = newSales.UUID
		if branchProduct.SellingPrice != nil {
			salesProduct.Price = *branchProduct.SellingPrice
		} else {
			salesProduct.Price = 0.0
		}

		_, errW = s.salesProductDomain.Create(ctx, salesProduct)
		if errW != nil {
			fmt.Println("Error creating sales product ", errW)
			continue
		}

		referenceType := constant.Sales
		for _, recipe := range productRecipes {
			errW = s.stockTransactionDomain.Create(model.StockTransaction{
				BranchOriginID:      payload.BranchID,
				BranchDestinationID: payload.BranchID,
				ItemID:              recipe.ItemID,
				Type:                "OUT",
				IssuerID:            userID,
				Quantity:            recipe.Amount * sales.Quantity,
				Cost:                salesProduct.Cost,
				Unit:                recipe.Unit,
				Reference:           salesProduct.UUID,
				ReferenceType:       &referenceType,
			})

			if errW != nil {
				fmt.Println("Error creating stock transction", errW)
				continue
			}

			errW = s.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
				ItemID:   recipe.ItemID,
				BranchID: payload.BranchID,
				NewTime:  payload.TransactionDate,
			})
			if errW != nil {
				fmt.Println("Error recalcualting inventory ", errW)
				continue
			}
		}

	}

	return nil

}

func (s *salesService) Update(ctx context.Context, payload dto.UpdateSalesRequest) (errW *error_wrapper.ErrorWrapper) {
	// 1. Get old sales data
	// 2.
	return
}

func (s *salesService) Delete(ctx context.Context, salesID string, userID string) *error_wrapper.ErrorWrapper {
	// First, get the sales data before deleting to create reversing transactions
	_, errW := s.salesDomain.FindByID(salesID)
	if errW != nil {
		return errW
	}

	_, errW = s.salesDomain.Delete(ctx, salesID)
	if errW != nil {
		return errW
	}

	_, errW = s.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": salesID,
		},
	}, userID)

	if errW != nil {
		fmt.Println("Error invalidating stock transaction", errW)
		return errW
	}

	return nil
}

func (s *salesService) Get(ctx context.Context, payload dto.GetListRequest) (sales []dto.GetSalesListResponse, errW *error_wrapper.ErrorWrapper) {
	return s.salesDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
}
