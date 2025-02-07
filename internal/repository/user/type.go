package user

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(name, username, email, password string, role model.UserRole) *error_wrapper.ErrorWrapper
	FindById(id string) (*model.User, *error_wrapper.ErrorWrapper)
	FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper)
}

type userRepository struct {
	db *gorm.DB
}
