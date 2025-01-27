package supplier

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
)

func (s *supplierController) GetSuppliers(c *gin.Context) {
	suppliers, err := s.supplierService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

func (s *supplierController) GetSupplier(c *gin.Context) {
	id := c.Param("id")
	supplier, err := s.supplierService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

func (s *supplierController) CreateSupplier(c *gin.Context) {
	var createSupplierRequest dto.CreateSupplierRequest
	if err := c.ShouldBindJSON(&createSupplierRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.supplierService.Create(createSupplierRequest.Name, createSupplierRequest.PhoneNumber, createSupplierRequest.Address, createSupplierRequest.PICName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Supplier created successfully"})
}

func (s *supplierController) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var updateSupplierRequest dto.UpdateSupplierRequest
	if err := c.ShouldBindJSON(&updateSupplierRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.supplierService.Update(id, updateSupplierRequest.Name, updateSupplierRequest.PhoneNumber, updateSupplierRequest.Address, updateSupplierRequest.PICName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated successfully"})
}

func (s *supplierController) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	err := s.supplierService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}
