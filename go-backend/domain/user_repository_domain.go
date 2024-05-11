package domain

import "context"

type User struct {
	ID uint	`gorm:"primaryKey"`
	Email string
	PasswordHash []byte
}

func (u1 User) Equals(u2 User) bool {
	return u1.ID == u2.ID && u1.Email == u2.Email
}

type UserRepository interface {
	CheckExistByEmail(ctx context.Context, email string) (bool, error)
	Create(context.Context, *User) error
	Fetch(ctx context.Context, conds *User) (*User, error)
}