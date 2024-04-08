package controller

import (
	"github.com/gin-gonic/gin"

	"go-backend/domain"
	"go-backend/setup"
)

type signupController struct {
	signupUsecase domain.SignupUsecase
	env *setup.Env
}

func NewSignupController(su domain.SignupUsecase, env *setup.Env) signupController {
	return signupController{
		signupUsecase: su,
		env: env,
	};
}

func (sc *signupController) Signup(*gin.Context) {

}