package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

func (s *salesController) Create(c *gin.Context) {
	var (
		errW    *error_wrapper.ErrorWrapper
		payload dto.CreateSalesRequest
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
		return
	}()

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = s.salesService.Create(c, payload)

	if errW != nil {
		return
	}

	return
}
