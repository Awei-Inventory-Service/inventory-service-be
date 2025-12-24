package sales

import (
	"context"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
	"github.com/inventory-service/resource/branch_product"
	"github.com/inventory-service/resource/product"
	"github.com/inventory-service/resource/sales"
	sales_product_resource "github.com/inventory-service/resource/sales_product"
)

type SalesDomain interface {
	Create(ctx context.Context, payload model.Sales) (*model.Sales, *error_wrapper.ErrorWrapper)
	FindAll() ([]model.Sales, *error_wrapper.ErrorWrapper)
	FindByID(id string) (*model.Sales, *error_wrapper.ErrorWrapper)
	Update(id string, sale model.Sales) *error_wrapper.ErrorWrapper
	Delete(ctx context.Context, id string) (*model.Sales, *error_wrapper.ErrorWrapper)
	Get(ctx context.Context, filter []dto.Filter, order []dto.Order, limit, offset int) ([]dto.GetSalesListResponse, *error_wrapper.ErrorWrapper)
}

type salesDomain struct {
	salesResource         sales.SalesResource
	productResource       product.ProductResource
	branchProductResource branch_product.BranchProductResource
	salesProductResource  sales_product_resource.SalesProductResource
}
