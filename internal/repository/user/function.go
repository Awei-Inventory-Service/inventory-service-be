package user

import (
	"github.com/inventory-service/internal/model"
)

func (r *userRepository) Create(name, username, email, role, password string) error {
	user := model.User{
		Name:     name,
		Username: username,
		Email:    email,
		Role:     role,
	}

	user.HashPassword(password)
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) FindUserByIdentifier(identifier string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) FindById(id string) (*model.User, error) {
	var user model.User
	result := r.db.Where("uuid = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
