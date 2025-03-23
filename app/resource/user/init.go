package user

import "gorm.io/gorm"

func NewUserResource(db *gorm.DB) UserResource {
	return &userResource{db: db}
}
