package supplier

import "github.com/inventory-service/app/Domain/supplier"

func NewSupplierService(supplierDomain supplier.SupplierDomain) SupplierService {
	return &supplierService{supplierDomain: supplierDomain}
}
