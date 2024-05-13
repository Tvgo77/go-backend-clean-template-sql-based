package controller

import (
	"go-backend/domain"
	"go-backend/setup"

	"github.com/gin-gonic/gin"
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

}

func (pc *profileController) UpdateProfile(c *gin.Context) {
	
}