package domain

import "context"

type FetchProfileRequest struct {
	UID int `json:"uid" binding:"required"`
}

type FetchProfileResponse struct {
	Profile Profile  `json:"profile"`
}

type UpdateProfileRequest struct {
	UID int `json:"uid" binding:"required"`
	Profile Profile  `json:"profile" binding:"required"`
}

type UpdateProfileResponse struct {
	Message string  `json:"message"`
}

type ProfileUsecase interface {
	GetUserByUID(ctx context.Context, UID int) (*User, error)
	UpdateProfile(ctx context.Context, user *User) error
}