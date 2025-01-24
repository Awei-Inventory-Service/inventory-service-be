package auth

import (
	"errors"

	"github.com/inventory-service/internal/utils"
)

type UserService interface {
	Login(identifier, password string) (string, error)
	Register(name, username, email, password string) error
	UpdateRole(username, role string) error
}

type userService struct {
	userRepository userRepository
}

func (u *userService) Login(identifier, password string) (string, error) {
	var token string

	user, err := u.userRepository.FindUserByIdentifier(identifier)
	if err != nil {
		return token, err
	}

	if !user.CheckPassword(password) {
		return token, errors.New("invalid credentials")
	}

	token, err = utils.GenerateToken(user.UUID, user.Name, user.Username, user.Email, user.Role)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (u *userService) Register(name, username, email, password string) error {
	errChan := make(chan error, 2)

	go func() {
		user, err := u.userRepository.FindUserByIdentifier(username)
		if err == nil && user != nil {
			errChan <- errors.New("username already exists")
		} else {
			errChan <- nil
		}
	}()

	go func() {
		emailUser, err := u.userRepository.FindUserByIdentifier(email)
		if err == nil && emailUser != nil {
			errChan <- errors.New("email already exists")
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
		return errors.New("username already exists")
	}

	// Check if email already exists
	existingEmail, err := u.userRepository.FindUserByIdentifier(email)
	if err == nil && existingEmail != nil {
		return errors.New("email already exists")
	}

	err = u.userRepository.Create(name, username, email, "Guest", password)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) UpdateRole(username, role string) error {
	panic("not implemented") // TODO: Implement
}
