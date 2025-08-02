package upload

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (u *uploadControllter) UploadTransaction(c *gin.Context) {
	var (
		errW    *error_wrapper.ErrorWrapper
		payload dto.UploadTransaction
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	// if err := c.ShouldBindJSON(&payload); err != nil {
	// 	errW = error_wrapper.New(model.CErrJsonBind, err.Error())
	// 	return
	// }

	branchId := c.GetHeader("branch_id")

	if branchId == "" {
		errW = error_wrapper.New(model.CErrHeaderIncomplete, "Branch id is required")
		return
	}
	payload.BranchID = branchId

	file, err := c.FormFile("file")
	if err != nil {
		errW = error_wrapper.New(model.CErrPayloadIncomplete, err.Error())
		return
	}

	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		errW = error_wrapper.New(model.CErrFileUpload, err.Error())
		return
	}

	errW = u.uploadService.ParseTransactionExcel(c, filePath, branchId)

	if errW != nil {
		return
	}

}
