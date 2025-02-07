package auth

import (
	"github.com/inventory-service/internal/model"
	"github.com/inventory-service/internal/utils"
	"github.com/inventory-service/lib/error_wrapper"
)

func (u *userService) Login(identifier, password string) (string, *error_wrapper.ErrorWrapper) {
	var token string

	user, errW := u.userRepository.FindUserByIdentifier(identifier)
	if errW != nil {
		return token, errW
	}

	if !user.CheckPassword(password) {
		return token, error_wrapper.New(model.SErrAuthInvalidCredentials, "Invalid credentials")
	}

	token, err := utils.GenerateToken(user.UUID, user.Name, user.Username, user.Email, string(user.Role))
	if err != nil {
		return token, error_wrapper.New(model.SErrAuthGenerateToken, "Error generating JWT token")
	}

	return token, nil
}

func (u *userService) Register(name, username, email, password string) *error_wrapper.ErrorWrapper {
	errChan := make(chan *error_wrapper.ErrorWrapper, 2)

	go func() {
		user, err := u.userRepository.FindUserByIdentifier(username)
		if err == nil && user != nil {
			errChan <- error_wrapper.New(model.SErrDataExist, "Username already exist")
		} else {
			errChan <- nil
		}
	}()

	go func() {
		emailUser, err := u.userRepository.FindUserByIdentifier(email)
		if err == nil && emailUser != nil {
			errChan <- error_wrapper.New(model.SErrDataExist, "Email already exist")
		} else {
			errChan <- nil
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	// Check if username already exists
	existingUser, err := u.userRepository.FindUserByIdentifier(username)
	if err == nil && existingUser != nil {
		return error_wrapper.New(model.SErrDataExist, "Username already exist")
	}

	// Check if email already exists
	existingEmail, err := u.userRepository.FindUserByIdentifier(email)
	if err == nil && existingEmail != nil {
		return error_wrapper.New(model.SErrDataExist, "Email already exist")
	}

	err = u.userRepository.Create(name, username, email, password, model.RoleGuest)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) UpdateRole(username, role string) *error_wrapper.ErrorWrapper {
	panic("not implemented") // TODO: Implement
}
