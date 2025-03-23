package user

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (u *userDomain) Create(name, username, email, password string, role model.UserRole) *error_wrapper.ErrorWrapper {
	user := model.User{
		Name:     name,
		Username: username,
		Email:    email,
		Role:     role,
	}

	user.HashPassword(password)

	return u.userResource.Create(user, role)
}

func (u *userDomain) FindById(id string) (*model.User, *error_wrapper.ErrorWrapper) {
	return u.userResource.FindById(id)
}

func (u *userDomain) FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper) {
	return u.userResource.FindUserByIdentifier(identifier)
}
