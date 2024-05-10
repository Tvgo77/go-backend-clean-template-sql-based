package domain

import (
	"context"

)

type LoginRequest struct {
	Email string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
}

type LoginUsecase interface {
	HasUser(ctx context.Context, email string) (bool, error)
	NewJWTtoken(user *User) (string, error)
}