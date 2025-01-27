package supplier

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/service/supplier"
)

type SupplierController interface {
	GetSuppliers(c *gin.Context)
	GetSupplier(c *gin.Context)
	CreateSupplier(c *gin.Context)
	UpdateSupplier(c *gin.Context)
	DeleteSupplier(c *gin.Context)
}

type supplierController struct {
	supplierService supplier.SupplierService
}
