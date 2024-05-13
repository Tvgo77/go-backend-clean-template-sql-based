package domain

type FetchProfileRequest struct {
	UID int `json:"uid" binding:"required"`
}

type FetchProfileResponse struct {
	Profile  // Embedded JSON struct
}

type UpdateProfileRequest struct {
	UID int `json:"uid" binding:"required"`
	Profile Profile  `json:"profile" binding:"required"`
}

type UpdateProfileResponse struct {
	Message string  `json:"message"`
}

type ProfileUsecase interface {
	
}