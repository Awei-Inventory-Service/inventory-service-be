package supplier

import "github.com/inventory-service/resource/supplier"

func NewSupplierDomain(supplierResource supplier.SupplierResource) SupplierDomain {
	return &supplierDomain{supplierResource: supplierResource}
}
