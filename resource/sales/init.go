package sales

import "gorm.io/gorm"

func NewSalesResource(db *gorm.DB) SalesResource {
	return &salesResource{db: db}
}
