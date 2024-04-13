package domain

import "context"

type User struct {
	ID uint	`gorm:"primaryKey"`
	Email string
	PasswordHash []byte
}

type UserRepository interface {
	CheckExistByEmail(ctx context.Context, email string) (bool, error)
	Create(context.Context, *User) error
}