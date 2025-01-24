package consignmentitem

import "gorm.io/gorm"

func NewConsignmentItemRepository(db *gorm.DB) ConsignmentItemRepository {
	return &consignmentItemRepository{db: db}
}
