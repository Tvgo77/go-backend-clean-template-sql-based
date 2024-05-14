package domain

import "context"

type FetchProfileRequest struct {
	
}

type FetchProfileResponse struct {
	Profile Profile  `json:"profile"`
}

type UpdateProfileRequest struct {
	Profile Profile  `json:"profile" binding:"required"`
}

type UpdateProfileResponse struct {
	Message string  `json:"message"`
}

type ProfileUsecase interface {
	GetUserByUID(ctx context.Context, UID uint) (*User, error)
	UpdateProfile(ctx context.Context, user *User) error
}