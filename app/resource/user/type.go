package user

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type UserResource interface {
	Create(user model.User, role model.UserRole) *error_wrapper.ErrorWrapper
	FindById(id string) (*model.User, *error_wrapper.ErrorWrapper)
	FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper)
}

type userResource struct {
	db *gorm.DB
}
