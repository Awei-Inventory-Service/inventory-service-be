package consignmentitem

import "gorm.io/gorm"

func NewConsignmentItemResource(db *gorm.DB) ConsignmentItemResource {
	return &consignmentItemResource{db: db}
}
