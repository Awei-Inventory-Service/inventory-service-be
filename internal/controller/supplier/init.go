package supplier

import "github.com/inventory-service/internal/service/supplier"

func NewSupplierController(supplierService supplier.SupplierService) SupplierController {
	return &supplierController{
		supplierService: supplierService,
	}
}
