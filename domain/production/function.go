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
		var (
			waste           float64
			wastePercentage float64
		)

		if payload.FinalUnit == productSourceItem.InitialUnit {
			waste = productSourceItem.InitialQuantity - payload.FinalQuantity

			wastePercentage = waste / productSourceItem.InitialQuantity * 100
		}

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

func (p *productionDomain) Get(ctx context.Context, filter dto.GetProductionFilter) ([]dto.GetProductionList, *error_wrapper.ErrorWrapper) {
	var (
		productionResults []dto.GetProductionList
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

func (p *productionDomain) mapToGetProduction(production model.Production) dto.GetProductionList {
	var (
		result dto.GetProductionList
	)

	result.BranchID = production.BranchID
	result.BranchName = production.Branch.Name
	result.FinalItemID = production.FinalItemID
	result.FinalItemName = production.FinalItem.Name
	result.FinalQuantity = production.FinalQuantity
	result.FinalUnit = production.FinalUnit
	result.ProductionDate = production.ProductionDate.String()
	result.ProductionID = production.UUID

	for _, productionItem := range production.SourceItems {
		result.SourceItems = append(result.SourceItems, dto.GetProductionItem{
			SourceItemID:    productionItem.UUID,
			SourceItemName:  productionItem.SourceItem.Name,
			InitialQuantity: productionItem.Quantity,
			Waste:           productionItem.WasteQuantity,
			WastePercentage: productionItem.WastePercentage,
		})
	}
	return result
}

func (p *productionDomain) Delete(ctx context.Context, productionID string) *error_wrapper.ErrorWrapper {
	errW := p.productionResource.Delete(ctx, productionID)

	if errW != nil {
		return errW
	}

	errW = p.productionItemResource.Delete(ctx, model.ProductionItem{
		ProductionID: productionID,
	})

	if errW != nil {
		return errW
	}
	return nil
}

func (p *productionDomain) Update(ctx context.Context, payload dto.UpdateProductionRequest) (errW *error_wrapper.ErrorWrapper) {
	productionDate, err := time.Parse("2006-01-02", payload.ProductionDate)
	if err != nil {
		return error_wrapper.New(model.ErrInvalidTimestamp, err.Error())
	}
	productionPayload := model.Production{
		FinalItemID:    payload.FinalItemID,
		FinalQuantity:  payload.FinalQuantity,
		FinalUnit:      payload.FinalUnit,
		BranchID:       payload.BranchID,
		ProductionDate: productionDate,
		UpdatedAt:      time.Now(),
	}

	return p.productionResource.Update(payload.ProductionID, productionPayload)
}
