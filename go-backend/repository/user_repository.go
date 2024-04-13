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

func (ur *userRepository) CheckExistByEmail(ctx context.Context, email string) (bool, error) {
	// Set timeout for database query
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ur.env.TimeoutSeconds))
	defer cancel()

	count, err := ur.database.Count(ctx, &domain.User{Email: email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}