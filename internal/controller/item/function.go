package item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
)

func (i *itemController) GetItems(c *gin.Context) {
	items, err := i.itemService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (i *itemController) GetItem(c *gin.Context) {
	id := c.Param("id")
	item, err := i.itemService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (i *itemController) CreateItem(c *gin.Context) {
	var createItemRequest dto.CreateItemRequest
	if err := c.ShouldBindJSON(&createItemRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.itemService.Create(createItemRequest.Name, createItemRequest.SupplierID, createItemRequest.Category, createItemRequest.Price, createItemRequest.Unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully"})
}

func (i *itemController) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updateItemRequest dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&updateItemRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := i.itemService.Update(id, updateItemRequest.SupplierID, updateItemRequest.Name, updateItemRequest.Category, updateItemRequest.Price, updateItemRequest.Unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func (i *itemController) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	err := i.itemService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
