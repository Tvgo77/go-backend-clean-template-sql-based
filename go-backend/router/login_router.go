package router

import (
	"github.com/gin-gonic/gin"

	"go-backend/controller"
	"go-backend/domain"
	"go-backend/repository"
	"go-backend/setup"
	"go-backend/usecase"
)

func LoginRouterSetup(env *setup.Env, db domain.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, env)
	lu := usecase.NewLoginUsecase(ur, env)
	lc := controller.NewLoginController(lu, env)

	group.Handle("POST", "/login", lc.Login)
}