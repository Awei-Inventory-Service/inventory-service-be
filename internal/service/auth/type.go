package auth

import (
	"github.com/inventory-service/internal/repository/user"
	"github.com/inventory-service/lib/error_wrapper"
)

type UserService interface {
	Login(identifier, password string) (string, *error_wrapper.ErrorWrapper)
	Register(name, username, email, password string) *error_wrapper.ErrorWrapper
	UpdateRole(username, role string) *error_wrapper.ErrorWrapper
}

type userService struct {
	userRepository user.UserRepository
}
