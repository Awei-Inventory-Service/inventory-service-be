package branch

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/dto"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
	"github.com/inventory-service/model"
)

func (b *branchController) GetBranches(c *gin.Context) {
	var (
		branches []model.Branch
		errW     *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, branches, errW)
	}()

	branches, errW = b.branchService.FindAll()
	if errW != nil {
		return
	}

}

func (b *branchController) GetBranch(c *gin.Context) {
	id := c.Param("id")
	var (
		branch *model.Branch
		errW   *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, branch, errW)
	}()
	branch, errW = b.branchService.FindByID(id)
	if errW != nil {
		return
	}

}

func (b *branchController) CreateBranch(c *gin.Context) {
	var (
		createBranchRequest dto.CreateBranchRequest
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createBranchRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = b.branchService.Create(createBranchRequest.Name, createBranchRequest.Location, createBranchRequest.BranchManagerID)
	if errW != nil {
		return
	}
}

func (b *branchController) UpdateBranch(c *gin.Context) {
	id := c.Param("id")
	var (
		updateBranchRequest dto.UpdateBranchRequest
		errW                *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()
	if err := c.ShouldBindJSON(&updateBranchRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = b.branchService.Update(id, updateBranchRequest.Name, updateBranchRequest.Location, updateBranchRequest.BranchManagerID)
	if errW != nil {
		return
	}
}

func (b *branchController) DeleteBranch(c *gin.Context) {
	var (
		errW *error_wrapper.ErrorWrapper
	)

	defer func() {
		response_wrapper.New(&c.Writer, c, errW == nil, nil, errW)
	}()
	id := c.Param("id")
	errW = b.branchService.Delete(id)
	if errW != nil {
		return
	}
}
