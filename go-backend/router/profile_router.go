package router

import (
	"github.com/gin-gonic/gin"

	"go-backend/controller"
	"go-backend/domain"
	"go-backend/repository"
	"go-backend/setup"
	"go-backend/usecase"
)

func ProfileRouterSetup(env *setup.Env, db domain.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, env)
	pu := usecase.NewProfileUsecase(ur, env)
	pc := controller.NewProfileController(pu, env)

	group.Handle("GET", "/profile/:uid", pc.FetchProfile)
	group.Handle("POST", "/profile/:uid", pc.UpdateProfile)
}