package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/middleware"
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
	hasUser, err := lu.userRepository.CheckExistByEmail(ctx, email)
	return hasUser, err
}

func (lu *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	conds := &domain.User{Email: email}
	user, err := lu.userRepository.Fetch(ctx, conds)
	return user, err
}

func (lu *loginUsecase) NewJWTtoken(user *domain.User) (string, error) {
	return middleware.NewJWTuidToken(user, lu.env.TokenSecret)
}