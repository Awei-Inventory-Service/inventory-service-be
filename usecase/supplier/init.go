package supplier

import "github.com/inventory-service/domain/supplier"

func NewSupplierService(supplierDomain supplier.SupplierDomain) SupplierService {
	return &supplierService{supplierDomain: supplierDomain}
}
