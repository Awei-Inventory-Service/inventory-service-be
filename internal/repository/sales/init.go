package sales

import "gorm.io/gorm"

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db: db}
}
