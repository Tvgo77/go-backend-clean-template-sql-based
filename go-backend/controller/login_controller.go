package controller

import (
	"go-backend/domain"
	"go-backend/setup"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	loginUsecase domain.LoginUsecase
	env *setup.Env
}

func NewLoginController(lu domain.LoginUsecase, env *setup.Env) *loginController {
	return &loginController{
		loginUsecase: lu,
		env: env,
	}
}

func (lc *loginController) Login(c *gin.Context) {
	// Check request format

	// Check if user exist

	// Generate new JWT access token

	// HTTP response
}