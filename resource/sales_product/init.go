package sales_product_resource

import "gorm.io/gorm"

func NewSalesProductResource(db *gorm.DB) SalesProductResource {
	return &salesProductResource{db: db}
}
