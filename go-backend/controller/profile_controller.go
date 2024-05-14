package controller

import (
	"fmt"
	"go-backend/domain"
	"go-backend/setup"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type profileController struct {
	profileUsecase domain.ProfileUsecase
	env *setup.Env
}

func NewProfileController(pu domain.ProfileUsecase, env *setup.Env) *profileController {
	return &profileController{
		profileUsecase: pu,
		env: env,
	}
}

func (pc *profileController) FetchProfile(c *gin.Context) {
	// Check request format
	var req domain.FetchProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Bad Request"})
		return
	}

	// Check if user exist
	// Fetch user
	user, err := pc.profileUsecase.GetUserByUID(c, req.UID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, &domain.ErrorResponse{Message: "User not found"})
			return
		}
		color.Red("%v", err)
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error"})
		return
	}

	// HTTP response
	c.JSON(http.StatusOK, &domain.FetchProfileResponse{Profile: user.Profile})
}

func (pc *profileController) UpdateProfile(c *gin.Context) {
	// Check request format
	var req domain.UpdateProfileRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &domain.ErrorResponse{Message: "Bad Request"})
		return
	}

	// Check if uid in request same with uid in JWT token
	if fmt.Sprint(req.UID) != c.GetString("userID") {
		c.JSON(http.StatusUnauthorized, &domain.ErrorResponse{Message: "Can only update your own profile"})
		return
	}

	// Update user
	user := domain.User{
		ID: uint(req.UID),
		Profile: req.Profile,
	}
	err = pc.profileUsecase.UpdateProfile(c, &user)
	if err != nil {
		color.Red("%v", err)
		c.JSON(http.StatusInternalServerError, &domain.ErrorResponse{Message: "Server Error"})
		return
	}

	// HTTP response
	c.JSON(http.StatusOK, &domain.UpdateProfileResponse{Message: "Profile update success"})
}