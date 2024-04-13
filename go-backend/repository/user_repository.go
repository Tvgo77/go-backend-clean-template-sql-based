package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
	"time"
)

type userRepository struct {
	database domain.Database
	env *setup.Env
}

func NewUserRepository(db domain.Database, env *setup.Env) domain.UserRepository {
	return &userRepository{database: db, env: env};
}

func (ur *userRepository) CheckExistByEmail(ctx context.Context, email string) bool {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ur.env.TimeoutSeconds))
	defer cancel()

	var user domain.User
	ur.database.WithContext(ctx).Select("id").Where(&domain.User{Email: email}).First(&user)
	return false
}