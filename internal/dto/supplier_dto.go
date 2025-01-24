package dto

type CreateSupplierRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PICName     string `json:"pic_name" binding:"required"`
}

type UpdateSupplierRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	PICName     string `json:"pic_name" binding:"required"`
}