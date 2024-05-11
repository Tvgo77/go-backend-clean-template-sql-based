package controller

import (
	"go-backend/domain"
	"go-backend/setup"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	var req domain.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	// Check if user exist
	hasUser, err := lc.loginUsecase.HasUser(c, req.Email)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !hasUser {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Email address or Password is wrong"})
		return
	}

	// Fetch user
	user, err := lc.loginUsecase.GetUserByEmail(c, req.Email)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Check Password
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Email address or Password is wrong"})
		return
	}

	// Generate new JWT access token
	token, err := lc.loginUsecase.NewJWTtoken(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// HTTP response
	resp := domain.LoginResponse{Message: "Login Success", Token: token}
	c.JSON(http.StatusOK, resp)
}