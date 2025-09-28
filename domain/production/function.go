package production

import (
	"context"
	"fmt"
	"time"

	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/model"
)

func (p *productionDomain) Create(ctx context.Context, payload dto.CreateProductionRequest) (*model.Production, *error_wrapper.ErrorWrapper) {

	productionDate, err := time.Parse("2006-01-02", payload.ProductionDate)
	if err != nil {
		return nil, error_wrapper.New(model.ErrInvalidTimestamp, err.Error())
	}

	production, errW := p.productionResource.Create(ctx, model.Production{
		FinalItemID:    payload.FinalItemID,
		FinalQuantity:  payload.FinalQuantity,
		BranchID:       payload.BranchID,
		FinalUnit:      payload.FinalUnit,
		ProductionDate: productionDate,
	})

	if errW != nil {
		fmt.Println("Error creating production", errW)
		return nil, errW
	}

	for _, productSourceItem := range payload.SourceItems {
		waste := productSourceItem.InitialQuantity - payload.FinalQuantity

		wastePercentage := waste / productSourceItem.InitialQuantity * 100
		_, errW := p.productionItemResource.Create(ctx, model.ProductionItem{
			ProductionID:    production.UUID,
			SourceItemID:    productSourceItem.SourceItemID,
			Quantity:        productSourceItem.InitialQuantity,
			Unit:            productSourceItem.InitialUnit,
			WasteQuantity:   waste,
			WastePercentage: wastePercentage,
		})

		if errW != nil {
			return nil, errW
		}
	}
	return production, nil
}

func (p *productionDomain) Get(ctx context.Context, filter model.Production) ([]dto.GetProduction, *error_wrapper.ErrorWrapper) {
	var (
		productionResults []dto.GetProduction
		errW              *error_wrapper.ErrorWrapper
	)
	productions, errW := p.productionResource.Get(ctx, filter)

	if errW != nil {
		return nil, errW
	}

	for _, production := range productions {
		mappedProduction := p.mapToGetProduction(production)
		productionResults = append(productionResults, mappedProduction)
	}

	return productionResults, nil
}

func (p *productionDomain) mapToGetProduction(production model.Production) dto.GetProduction {
	return dto.GetProduction{
		UUID:           production.UUID,
		FinalItemID:    production.FinalItemID,
		FinalItemName:  production.FinalItem.Name,
		BranchID:       production.BranchID,
		BranchName:     production.Branch.Name,
		ProductionDate: production.ProductionDate.Format("2006-01-02"),
	}
}
