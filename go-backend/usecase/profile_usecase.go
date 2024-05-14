package usecase

import (
	"context"
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

func (pu *profileUsecase) GetUserByUID(ctx context.Context, UID int) (*domain.User, error) {
	return nil, nil
}