package domain

import (
	"context"
	"time"
)

type User struct {
	ID uint	`gorm:"primaryKey"`
	Email string
	PasswordHash []byte
	Profile Profile  `gorm:"embedded"`
}

type Profile struct {
	Name string  `json:"name"`
	Bio string  `json:"bio"`
	BirthDay time.Time `gorm:"type:date"`
}

func (u1 User) Equals(u2 User) bool {
	return u1.ID == u2.ID && u1.Email == u2.Email
}

type UserRepository interface {
	CheckExistByEmail(ctx context.Context, email string) (bool, error)
	Create(context.Context, *User) error
	Fetch(ctx context.Context, conds *User) (*User, error)
}