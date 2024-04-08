package usecase

import (
	"go-backend/domain"
	"go-backend/setup"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env 
}

func NewSignupUsecase(ur domain.UserRepository, env *setup.Env) signupUsecase {
	return signupUsecase{
		userRepository: ur,
		env: env,
	}
}

