package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-backend/domain"
	"go-backend/setup"
)

type signupController struct {
	signupUsecase domain.SignupUsecase
	env *setup.Env
}

func NewSignupController(su domain.SignupUsecase, env *setup.Env) *signupController {
	return &signupController{
		signupUsecase: su,
		env: env,
	}
}

func (sc *signupController) Signup(c *gin.Context) {
	// Check request format
	var request domain.SignupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if register email already exist
	hasUser, err := sc.signupUsecase.HasUser(c, request.Email)
	if err != nil {
		log.Fatal(err)
		return
	}

	if hasUser {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: ""})
		return
	}
	
	// Create new user

	// Response success
}