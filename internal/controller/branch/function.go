package branch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"github.com/inventory-service/lib/response_wrapper"
)

func (b *branchController) GetBranches(c *gin.Context) {
	var branches []model.Branch
	branches, err := b.branchService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branches)
}

func (b *branchController) GetBranch(c *gin.Context) {
	id := c.Param("id")
	branch, err := b.branchService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, branch)
}

func (b *branchController) CreateBranch(c *gin.Context) {
	var (
		createBranchRequest dto.CreateBranchRequest
		errW                *error_wrapper.ErrorWrapper
		isSuccess           bool
	)

	defer func() {
		if errW == nil {
			isSuccess = true
		} else {
			isSuccess = false
		}
		response_wrapper.New(&c.Writer, c, isSuccess, nil, errW)
	}()

	if err := c.ShouldBindJSON(&createBranchRequest); err != nil {
		errW = error_wrapper.New(model.CErrJsonBind, err.Error())
		return
	}

	errW = b.branchService.Create(createBranchRequest.Name, createBranchRequest.Location, createBranchRequest.BranchManagerID)
	if errW != nil {
		return
	}

	isSuccess = true
}

func (b *branchController) UpdateBranch(c *gin.Context) {
	id := c.Param("id")
	var updateBranchRequest dto.UpdateBranchRequest
	if err := c.ShouldBindJSON(&updateBranchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := b.branchService.Update(id, updateBranchRequest.Name, updateBranchRequest.Location, updateBranchRequest.BranchManagerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (b *branchController) DeleteBranch(c *gin.Context) {
	id := c.Param("id")
	err := b.branchService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
