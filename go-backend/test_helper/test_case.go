package test_helper

import (
	"go-backend/domain"
)

var SignupReqBody = domain.SignupRequest{
	Email:    "test@gmail.com",
	Password: "password",
}

var LoginReqBody = domain.LoginRequest{
	Email:    "test@gmail.com",
	Password: "password",
}

var TestUser = domain.User{
	ID: 1,
	Email: "test@gmail.com",
	PasswordHash: []byte("$2a$08$2yWRafKKuOWV.9A1dHbpqOughYDzyi8ZqrXC.i4dbWq3/YNxTzIw."),  // Precomputed
}

