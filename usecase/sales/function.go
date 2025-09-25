package sales

import (
	"context"
	"time"

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

	// Set transaction date from request, or use current time if not provided
	if payload.TransactionDate != "" {
		parsedDate, err := time.Parse("2006-01-02", payload.TransactionDate)
		if err != nil {
			return error_wrapper.New(model.CErrJsonBind, "Invalid date format. Expected YYYY-MM-DD")
		}
		sales.TransactionDate = parsedDate
	} else {
		sales.TransactionDate = time.Now()
	}

	_, errW = s.salesDomain.Create(ctx, sales)

	if errW != nil {
		return errW
	}

	for _, itemComposition := range product.ProductComposition {
		total := itemComposition.Ratio * payload.Quantity
		referenceType := "SALES_CREATION"
		errW = s.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      payload.BranchID,
			BranchDestinationID: payload.BranchID,
			ItemID:              itemComposition.ItemID,
			Type:                "OUT",
			IssuerID:            userID,
			Quantity:            total,
			Cost:                sales.Cost,
			Unit:                itemComposition.Item.Unit,
			Reference:           "", // Will be updated with sales ID after creation
			ReferenceType:       &referenceType,
		})

		if errW != nil {
			return errW
		}
		errW = s.branchItemDomain.SyncBranchItem(ctx, payload.BranchID, itemComposition.ItemID)

		if errW != nil {
			return errW
		}

	}

	return nil
}

func (s *salesService) Delete(ctx context.Context, salesID string, userID string) *error_wrapper.ErrorWrapper {
	// First, get the sales data before deleting to create reversing transactions
	salesData, errW := s.salesDomain.FindByID(salesID)
	if errW != nil {
		return errW
	}

	// Get product composition to reverse stock transactions
	product, errW := s.productDomain.FindByID(ctx, salesData.BranchProduct.ProductID)
	if errW != nil {
		return errW
	}

	_, errW = s.salesDomain.Delete(ctx, salesID)
	if errW != nil {
		return errW
	}

	for _, itemComposition := range product.ProductComposition {
		total := itemComposition.Ratio * salesData.Quantity
		referenceType := "SALES_DELETION"
		errW = s.stockTransactionDomain.Create(model.StockTransaction{
			BranchOriginID:      salesData.BranchProduct.BranchID,
			BranchDestinationID: salesData.BranchProduct.BranchID,
			ItemID:              itemComposition.ItemID,
			Type:                "IN", // Opposite of original "OUT" transaction
			IssuerID:            userID,
			Quantity:            total,
			Cost:                salesData.Cost,
			Unit:                itemComposition.Item.Unit,
			Reference:           salesID,
			ReferenceType:       &referenceType,
		})

		if errW != nil {
			return errW
		}
	}

	return nil
}

func (s *salesService) FindGroupedByDate(ctx context.Context) ([]dto.SalesGroupedByDateResponse, *error_wrapper.ErrorWrapper) {
	sales, errW := s.salesDomain.FindGroupedByDate(ctx)
	if errW != nil {
		return nil, errW
	}

	groupedSales := make(map[string][]model.Sales)
	for _, sale := range sales {
		dateKey := sale.TransactionDate.Format("2006-01-02")
		groupedSales[dateKey] = append(groupedSales[dateKey], sale)
	}

	var response []dto.SalesGroupedByDateResponse
	for date, salesList := range groupedSales {
		var totalRevenue, totalProfit float64
		var salesResponses []dto.GetSalesResponse

		for _, sale := range salesList {
			totalRevenue += sale.Price * sale.Quantity
			totalProfit += (sale.Price - sale.Cost) * sale.Quantity

			salesResponses = append(salesResponses, dto.GetSalesResponse{
				BranchID:    sale.BranchProduct.BranchID,
				BranchName:  sale.BranchProduct.Branch.Name,
				ProductID:   sale.BranchProduct.ProductID,
				ProductName: sale.BranchProduct.Product.Name,
				Quantity:    sale.Quantity,
			})
		}

		response = append(response, dto.SalesGroupedByDateResponse{
			TransactionDate: date,
			TotalSales:      len(salesList),
			TotalRevenue:    totalRevenue,
			TotalProfit:     totalProfit,
			Sales:           salesResponses,
		})
	}

	return response, nil
}

func (s *salesService) FindGroupedByDateAndBranch(ctx context.Context) ([]dto.SalesGroupedByDateAndBranchResponse, *error_wrapper.ErrorWrapper) {
	sales, errW := s.salesDomain.FindGroupedByDateAndBranch(ctx)
	if errW != nil {
		return nil, errW
	}

	// Group sales by date+branch combination (each group becomes one response object)
	dateBranchGroups := make(map[string][]model.Sales)
	for _, sale := range sales {
		dateKey := sale.TransactionDate.Format("2006-01-02")
		branchKey := sale.BranchProduct.BranchID
		// Create composite key: date + branch
		compositeKey := dateKey + "_" + branchKey
		dateBranchGroups[compositeKey] = append(dateBranchGroups[compositeKey], sale)
	}

	// Convert to response format - one object per date+branch combination
	var response []dto.SalesGroupedByDateAndBranchResponse
	for _, salesList := range dateBranchGroups {
		if len(salesList) == 0 {
			continue
		}

		var totalRevenue, totalProfit float64
		var salesResponses []dto.GetSalesResponse

		// Get date and branch info from first sale (all sales in group have same date/branch)
		firstSale := salesList[0]
		dateKey := firstSale.TransactionDate.Format("2006-01-02")
		branchID := firstSale.BranchProduct.BranchID
		branchName := firstSale.BranchProduct.Branch.Name

		// Process all sales in this date+branch group
		for _, sale := range salesList {
			totalRevenue += sale.Price * sale.Quantity
			totalProfit += (sale.Price - sale.Cost) * sale.Quantity

			salesResponses = append(salesResponses, dto.GetSalesResponse{
				BranchID:    sale.BranchProduct.BranchID,
				BranchName:  sale.BranchProduct.Branch.Name,
				ProductID:   sale.BranchProduct.ProductID,
				ProductName: sale.BranchProduct.Product.Name,
				Quantity:    sale.Quantity,
			})
		}

		response = append(response, dto.SalesGroupedByDateAndBranchResponse{
			TransactionDate: dateKey,
			TotalSales:      len(salesList),
			TotalRevenue:    totalRevenue,
			TotalProfit:     totalProfit,
			BranchID:        branchID,
			BranchName:      branchName,
			Sales:           salesResponses,
		})
	}

	return response, nil
}
