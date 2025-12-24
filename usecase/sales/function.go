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
				Reference:           newSales.UUID,
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
	oldSales, errW := s.salesDomain.Get(ctx, []dto.Filter{{
		Key:    "uuid",
		Values: []string{payload.SalesID},
	}}, nil, 0, 0)

	if errW != nil {
		fmt.Println("Error getting sales domain ", errW)
		return
	}

	deletedSales := oldSales[0]

	// 2. Delete old sales product data
	errW = s.salesProductDomain.Delete(ctx, model.SalesProduct{
		SalesID: deletedSales.SalesID,
	})
	if errW != nil {
		fmt.Println("Error deleting sales product domain ", errW)
		return
	}

	// 3. Invalidate stock transaction data
	deletedItems, errW := s.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": deletedSales.SalesID,
		},
	}, payload.UserID)
	if errW != nil {
		fmt.Println("Error invalidating stock transaction", errW)
		return
	}

	// 4. Delete old sales data
	_, errW = s.salesDomain.Delete(ctx, deletedSales.SalesID)
	if errW != nil {
		fmt.Println("Error deleting sales ", errW)
		return
	}

	parsedTransactionDate, err := time.Parse("2006-01-02", payload.TransactionDate)
	if err != nil {
		return error_wrapper.New(model.CErrJsonBind, "Invalid date format. Expected YYYY-MM-DD")
	}

	// 3. Create new sales data
	newSales := model.Sales{
		BranchID:        payload.BranchID,
		TransactionDate: parsedTransactionDate,
		UpdatedAt:       time.Now(),
	}

	errW = s.salesDomain.Update(payload.SalesID, newSales)
	if errW != nil {
		fmt.Println("Error updating sales record ", errW)
		return
	}

	for _, salesData := range payload.SalesData {
		var (
			salesProduct model.SalesProduct
		)

		product, errW := s.productDomain.FindByID(ctx, salesData.ProductID)
		if errW != nil {
			fmt.Println("Error finding product by id ", errW)
			continue
		}

		productRecipes, totalPrice, errW := s.productDomain.CalculateProductCost(ctx, *product, payload.BranchID, parsedTransactionDate)
		if errW != nil {
			fmt.Println("Error calculating product cost ", errW)
			return errW
		}

		branchProduct, errW := s.branchProductDomain.GetByBranchIdAndProductId(ctx, payload.BranchID, salesData.ProductID)
		if errW != nil {
			fmt.Println("Error getting branch product", errW)
			return errW
		}

		salesProduct.BranchID = payload.BranchID
		salesProduct.Cost = totalPrice * salesData.Quantity
		salesProduct.ProductID = product.UUID
		salesProduct.Quantity = salesData.Quantity
		salesProduct.Type = salesData.Type
		salesProduct.SalesID = payload.SalesID
		if branchProduct.SellingPrice != nil {
			salesProduct.Price = *branchProduct.SellingPrice
		} else {
			salesProduct.Price = 0.0
		}

		_, errW = s.salesProductDomain.Create(ctx, salesProduct)
		if errW != nil {
			fmt.Println("Error creating new sales product ", errW)
			return errW
		}

		referenceType := constant.Sales
		for _, recipe := range productRecipes {
			errW = s.stockTransactionDomain.Create(model.StockTransaction{
				BranchOriginID:      payload.BranchID,
				BranchDestinationID: payload.BranchID,
				ItemID:              recipe.ItemID,
				Type:                "OUT",
				IssuerID:            payload.UserID,
				Quantity:            recipe.Amount * salesData.Quantity,
				Cost:                salesProduct.Cost,
				Unit:                recipe.Unit,
				Reference:           newSales.UUID,
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

	parsedOldTransactionDate, err := time.Parse("2006-01-02", deletedSales.TransactionDate)
	if err != nil {
		return error_wrapper.New(model.CErrJsonBind, "Invalid date format. Expected YYYY-MM-DD")
	}
	for _, deletedItem := range deletedItems {
		errW = s.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			ItemID:       deletedItem,
			BranchID:     payload.BranchID,
			NewTime:      payload.TransactionDate,
			PreviousTime: &parsedOldTransactionDate,
		})
		if errW != nil {
			fmt.Println("Error recalcualting inventory ", errW)
			continue
		}
	}

	return
}

func (s *salesService) Delete(ctx context.Context, salesID string, userID string) *error_wrapper.ErrorWrapper {
	// First, get the sales data before deleting to create reversing transactions
	sales, errW := s.salesDomain.FindByID(salesID)
	if errW != nil {
		return errW
	}

	errW = s.salesProductDomain.Delete(ctx, model.SalesProduct{
		SalesID: salesID,
	})
	if errW != nil {
		fmt.Println("Error deleting sales product ", errW)
		return errW
	}
	deletedItems, errW := s.stockTransactionDomain.InvalidateStockTransaction(ctx, []map[string]interface{}{
		{
			"field": "reference",
			"value": sales.UUID,
		},
	}, userID)

	for _, item := range deletedItems {
		errW = s.inventoryDomain.RecalculateInventory(ctx, dto.RecalculateInventoryRequest{
			ItemID:   item,
			BranchID: sales.BranchID,
			NewTime:  sales.TransactionDate.Format("2006-01-02"),
		})
		if errW != nil {
			fmt.Println("Error recalculating inventory ", errW)
			return errW
		}
	}

	_, errW = s.salesDomain.Delete(ctx, salesID)
	if errW != nil {
		return errW
	}

	return nil
}

func (s *salesService) Get(ctx context.Context, payload dto.GetListRequest) (sales []dto.GetSalesListResponse, errW *error_wrapper.ErrorWrapper) {
	return s.salesDomain.Get(ctx, payload.Filter, payload.Order, payload.Limit, payload.Offset)
}
