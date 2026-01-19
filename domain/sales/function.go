package sales

import (
	"context"
	"fmt"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesDomain) Create(ctx context.Context, payload model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.Create(payload)
}

func (s *salesDomain) FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindAll()
}

func (s *salesDomain) FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindByID(id)
}

func (s *salesDomain) Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper {
	return s.salesResource.Update(id, sale)
}

func (s *salesDomain) Delete(ctx context.Context, id string) (*model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.Delete(ctx, id)
}

func (s *salesDomain) Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]dto.GetSalesListResponse, *error_wrapper.ErrorWrapper) {
	var (
		salesResponse []dto.GetSalesListResponse
	)
	sales, errW := s.salesResource.Get(ctx, filter, order, limit, offset)
	if errW != nil {
		fmt.Println("Error getting sales ", errW)
		return nil, errW
	}

	for _, rawSales := range sales {
		var (
			salesProduct []dto.GetSalesProductResponse
		)

		for _, rawSalesProduct := range rawSales.SalesProducts {
			salesProduct = append(salesProduct, dto.GetSalesProductResponse{
				ProductID:   rawSalesProduct.ProductID,
				ProductName: rawSalesProduct.Product.Name,
				Quantity:    rawSalesProduct.Quantity,
				Cost:        rawSalesProduct.Cost,
				Price:       rawSalesProduct.Price,
				Type:        rawSalesProduct.Type,
			})
		}

		salesResponse = append(salesResponse, dto.GetSalesListResponse{
			SalesID:         rawSales.UUID,
			BranchID:        rawSales.BranchID,
			BranchName:      rawSales.Branch.Name,
			TransactionDate: rawSales.TransactionDate.Format("2006-01-02"),
			SalesProducts:   salesProduct,
		})
	}
	return salesResponse, nil
}
