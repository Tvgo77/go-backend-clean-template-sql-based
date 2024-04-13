package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env 
}

func NewSignupUsecase(ur domain.UserRepository, env *setup.Env) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: ur,
		env: env,
	}
}

func (su *signupUsecase) HasUser(ctx context.Context, email string) bool {
	hasUser := su.userRepository.CheckExistByEmail(ctx, email)
	return hasUser
}

