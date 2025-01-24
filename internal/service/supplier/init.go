package supplier

import "github.com/inventory-service/internal/repository/supplier"

func NewSupplierService(supplierRepository supplier.SupplierRepository) SupplierService {
	return &supplierService{supplierRepository: supplierRepository}
}
