package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inventory-service/internal/dto"
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/service/branch"
)

type BranchController interface {
	GetBranches(c *gin.Context)
	GetBranch(c *gin.Context)
	CreateBranch(c *gin.Context)
	UpdateBranch(c *gin.Context)
	DeleteBranch(c *gin.Context)
}

type branchController struct {
	branchService branch.BranchService
}

func NewBranchController(branchService branch.BranchService) BranchController {
	return &branchController{
		branchService: branchService,
	}
}

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
	var createBranchRequest dto.CreateBranchRequest
	if err := c.ShouldBindJSON(&createBranchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := b.branchService.Create(createBranchRequest.Name, createBranchRequest.Location, createBranchRequest.BranchManagerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
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
