package domain

import "context"

type User struct {
	ID uint	`gorm:"primaryKey"`
	Email string
	Password string
}

type UserRepository interface {
	CheckExistByEmail(ctx context.Context, email string) bool
}