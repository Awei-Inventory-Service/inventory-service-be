package production

import "github.com/inventory-service/usecase/production"

func NewProductionHandler(productionUsecase production.ProductionUsecase) ProductionHandler {
	return &productionHandler{
		productionUsecase: productionUsecase,
	}
}
