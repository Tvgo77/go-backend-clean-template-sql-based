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

func (pu *profileUsecase) GetUserByUID(ctx context.Context, UID uint) (*domain.User, error) {
	conds := &domain.User{ID: UID}
	user, err := pu.userRepository.Fetch(ctx, conds)
	return user, err
}

func (pu *profileUsecase) UpdateProfile(ctx context.Context, user *domain.User) (error) {
	old := &domain.User{ID: user.ID}
	new := user
	err := pu.userRepository.Update(ctx, old, new)
	return err
}