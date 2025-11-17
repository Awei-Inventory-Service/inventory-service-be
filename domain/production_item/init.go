package production_item_domain

import productionitem "github.com/inventory-service/resource/production_item"

func NewProductionItemDomain(
	productionItemResource productionitem.ProductionItemResource,
) ProductionItemDomain {
	return &productionItemDomain{
		productionItemResource: productionItemResource,
	}
}
