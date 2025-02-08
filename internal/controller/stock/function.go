package stock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *stockController) GetStockByItemID(c *gin.Context) {
	id := c.Param("id")

	stock, err := s.stockService.GetStockByItemID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}
