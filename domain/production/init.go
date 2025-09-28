package production

import (
	"github.com/inventory-service/resource/production"
	productionitem "github.com/inventory-service/resource/production_item"
)

func NewProductionDomain(
	productionResource production.ProductionResource,
	productionItemResource productionitem.ProductionItemResource,
) ProductionDomain {
	return &productionDomain{
		productionResource:     productionResource,
		productionItemResource: productionItemResource,
	}
}
