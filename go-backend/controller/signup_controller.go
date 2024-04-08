package controller

import (
	"go-backend/domain"
	"go-backend/setup"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env *setup.Env
}