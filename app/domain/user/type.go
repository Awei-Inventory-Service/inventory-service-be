package user

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/app/resource/user"
	"github.com/inventory-service/lib/error_wrapper"
)

type UserDomain interface {
	Create(name, username, email, password string, role model.UserRole) *error_wrapper.ErrorWrapper
	FindById(id string) (*model.User, *error_wrapper.ErrorWrapper)
	FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper)
}

type userDomain struct {
	userResource user.UserResource
}
