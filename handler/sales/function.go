package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (s *salesController) Create(c *gin.Context) {
	var (
		errW    *error_wrapper.ErrorWrapper
		payload dto.CreateSalesRequest
	)

	defer func() {
		if r := recover(); r != nil {
			errW = error_wrapper.New(model.CErrInternalServer, "Internal server error")
		}
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&payload); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	userID := c.GetHeader("user_id")

	if userID == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "User id is required in header")
		return
	}

	errW = s.salesService.Create(c, payload, userID)

	if errW != nil {
		return
	}
}
