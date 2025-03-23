package dto

type CreateBranchRequest struct {
	Name string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	BranchManagerID string `json:"branch_manager_id" binding:"required"`
}

type UpdateBranchRequest struct {
	Name string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	BranchManagerID string `json:"branch_manager_id" binding:"required"`
}