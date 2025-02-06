package auth

import "github.com/inventory-service/internal/repository/user"

func NewUserService(userRepository user.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
