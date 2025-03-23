package consignmentitem

import consignmentitem "github.com/inventory-service/app/resource/consignment_item"

func NewConsignmentItemDomain(consignmentItemResource consignmentitem.ConsignmentItemResource) ConsignmentItemDomain {
	return &consignmentItemDomain{consignmentItemResource: consignmentItemResource}
}
