package auth

import "github.com/inventory-service/app/usecase/auth"

func NewAuthController(authService auth.UserService) AuthController {
	return &authController{
		authService: authService,
	}
}
