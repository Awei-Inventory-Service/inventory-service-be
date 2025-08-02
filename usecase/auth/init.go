package auth

import "github.com/inventory-service/domain/user"

func NewUserService(userDomain user.UserDomain) UserService {
	return &userService{
		userDomain: userDomain,
	}
}
