package domain

import (
	"context"
)

type SignupRequest struct {
	Email string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type SignupUsecase interface {
	HasUser(ctx context.Context, email string) (bool, error)
	CreateNewUser(ctx context.Context, user *User) error
}