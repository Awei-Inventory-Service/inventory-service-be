package user

import "github.com/inventory-service/app/resource/user"

func NewUserDomain(userResource user.UserResource) UserDomain {
	return &userDomain{userResource: userResource}
}
