package sales

import (
	"context"

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

func (s *salesDomain) FindGroupedByDate(ctx context.Context) ([]model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindGroupedByDate(ctx)
}

func (s *salesDomain) FindGroupedByDateAndBranch(ctx context.Context) ([]model.Sales, *error_wrapper.ErrorWrapper) {
	return s.salesResource.FindGroupedByDateAndBranch(ctx)
}
