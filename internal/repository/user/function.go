package user

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (r *userRepository) Create(name, username, email, password string, role model.UserRole) *error_wrapper.ErrorWrapper {
	user := model.User{
		Name:     name,
		Username: username,
		Email:    email,
		Role:     role,
	}

	user.HashPassword(password)
	result := r.db.Create(&user)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (r *userRepository) FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper) {
	var user model.User
	result := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &user, nil
}

func (r *userRepository) FindById(id string) (*model.User, *error_wrapper.ErrorWrapper) {
	var user model.User
	result := r.db.Where("uuid = ?", id).First(&user)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &user, nil
}
