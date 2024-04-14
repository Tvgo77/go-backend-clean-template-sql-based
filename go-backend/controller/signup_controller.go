package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Email already registered"})
		return
	}
	
	// Create new user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
	if err != nil {
		log.Fatal(err)
		return
	}

	user := domain.User{
		Email: request.Email, 
		PasswordHash: passwordHash,
	}
	err = sc.signupUsecase.CreateNewUser(c, &user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Generate JWT token
	token, err := sc.signupUsecase.NewJWTtoken(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Response success
	c.JSON(http.StatusOK, domain.SignupResponse{
		Message: "Account create success",
		Token: token,
	})
}