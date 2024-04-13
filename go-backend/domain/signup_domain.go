package domain

import "context"

type SignupRequest struct {
	Email string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type SignupResponse struct {

}

type SignupUsecase interface {
	HasUser(ctx context.Context, email string) bool
}