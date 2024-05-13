package usecase

import (
	"go-backend/domain"
	"go-backend/setup"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env
}

func NewProfileUsecase(ur domain.UserRepository, env *setup.Env) *profileUsecase {
	return &profileUsecase{
		userRepository: ur,
		env: env,
	}
}