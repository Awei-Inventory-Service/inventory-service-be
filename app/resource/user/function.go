package user

import (
	"github.com/inventory-service/app/model"
	"github.com/inventory-service/lib/error_wrapper"
)

func (r *userResource) Create(user model.User, role model.UserRole) *error_wrapper.ErrorWrapper {

	result := r.db.Create(&user)
	if result.Error != nil {
		return error_wrapper.New(model.RErrPostgresCreateDocument, result.Error.Error())
	}

	return nil
}

func (r *userResource) FindUserByIdentifier(identifier string) (*model.User, *error_wrapper.ErrorWrapper) {
	var user model.User
	result := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &user, nil
}

func (r *userResource) FindById(id string) (*model.User, *error_wrapper.ErrorWrapper) {
	var user model.User
	result := r.db.Where("uuid = ?", id).First(&user)
	if result.Error != nil {
		return nil, error_wrapper.New(model.RErrPostgresReadDocument, result.Error.Error())
	}

	return &user, nil
}
