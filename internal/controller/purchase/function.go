package purchase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
)

func (p *purchaseController) GetPurchases(c *gin.Context) {
	purchases, err := p.purchaseService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchases)
}

func (p *purchaseController) GetPurchase(c *gin.Context) {
	id := c.Param("id")
	purchase, err := p.purchaseService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func (p *purchaseController) CreatePurchase(c *gin.Context) {
	var createPurchaseRequest dto.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&createPurchaseRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.purchaseService.Create(
		createPurchaseRequest.SupplierID,
		createPurchaseRequest.BranchID,
		createPurchaseRequest.ItemID,
		createPurchaseRequest.Quantity,
		createPurchaseRequest.PurchaseCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Purchase created successfully"})
}

func (p *purchaseController) UpdatePurchase(c *gin.Context) {
	id := c.Param("id")
	var updatePurchaseRequest dto.UpdatePurchaseRequest
	if err := c.ShouldBindJSON(&updatePurchaseRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.purchaseService.Update(
		id,
		updatePurchaseRequest.SupplierID,
		updatePurchaseRequest.BranchID,
		updatePurchaseRequest.ItemID,
		updatePurchaseRequest.Quantity,
		updatePurchaseRequest.PurchaseCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase updated successfully"})
}

func (p *purchaseController) DeletePurchase(c *gin.Context) {
	id := c.Param("id")
	err := p.purchaseService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase deleted successfully"})
}
