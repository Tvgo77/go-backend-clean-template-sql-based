package repository

import "go-backend/domain"

type userRepository struct {
	database domain.Database
}

func NewUserRepository(db domain.Database) userRepository {
	return userRepository{database: db};
}