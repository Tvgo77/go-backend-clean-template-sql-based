package router

import (
	"github.com/gin-gonic/gin"

	"go-backend/controller"
	"go-backend/domain"
	"go-backend/repository"
	"go-backend/setup"
	"go-backend/usecase"
)

func SignupRouterSetup(env *setup.Env, db domain.Database, group *gin.RouterGroup) {
	// Create controller, usecase, repository instance
	ur := repository.NewUserRepository(db, env)
	su := usecase.NewSignupUsecase(ur, env)
	sc := controller.NewSignupController(su, env)

	// Setup gin router
	group.Handle("POST", "/signup", sc.Signup)
}