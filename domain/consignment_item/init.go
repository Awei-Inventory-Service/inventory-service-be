package consignmentitem

import consignmentitem "github.com/inventory-service/resource/consignment_item"

func NewConsignmentItemDomain(consignmentItemResource consignmentitem.ConsignmentItemResource) ConsignmentItemDomain {
	return &consignmentItemDomain{consignmentItemResource: consignmentItemResource}
}
