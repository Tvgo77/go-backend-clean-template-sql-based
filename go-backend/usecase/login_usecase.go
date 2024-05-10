package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env
}

func NewLoginUsecase(ur domain.UserRepository, env *setup.Env) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		env: env,
	}
}

func (lu *loginUsecase) HasUser(ctx context.Context, email string) (bool, error) {
	return false, nil
}

func (lu *loginUsecase) NewJWTtoken(user *domain.User) (string, error) {
	return "", nil
}