package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/middleware"
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

func (su *signupUsecase) HasUser(ctx context.Context, email string) (bool, error) {
	hasUser, err := su.userRepository.CheckExistByEmail(ctx, email)
	return hasUser, err
}

func (su *signupUsecase) CreateNewUser(ctx context.Context, user *domain.User) error {
	err := su.userRepository.Create(ctx, user)
	return err
}

func (su *signupUsecase) NewJWTtoken(user *domain.User) (string, error) {
	return middleware.NewJWTuidToken(user, su.env.TokenSecret)
}
