package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/inventory-service/app/usecase/auth"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService auth.UserService
}
