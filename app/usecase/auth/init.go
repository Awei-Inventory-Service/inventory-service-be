package auth

import "github.com/inventory-service/app/domain/user"


func NewUserService(userDomain user.UserDomain) UserService {
	return &userService{
		userDomain: userDomain,
	}
}
