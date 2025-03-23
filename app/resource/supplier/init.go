package supplier

import "gorm.io/gorm"

func NewSupplierResource(db *gorm.DB) SupplierResource {
	return &supplierResource{db: db}
}
