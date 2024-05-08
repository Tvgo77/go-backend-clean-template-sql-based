package usecase

import (
	"context"
	"fmt"
	"go-backend/domain"
	"go-backend/setup"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env 
}

func NewSignupUsecase(ur domain.UserRepository, env *setup.Env) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: ur,
		env: env,
	}
}

func (su *signupUsecase) HasUser(ctx context.Context, email string) (bool, error) {
	hasUser, err := su.userRepository.CheckExistByEmail(ctx, email)
	return hasUser, err
}

func (su *signupUsecase) CreateNewUser(ctx context.Context, user *domain.User) error {
	err := su.userRepository.Create(ctx, user)
	return err
}

func (su *signupUsecase) NewJWTtoken(user *domain.User) (string, error) {
	myClaims := jwt.RegisteredClaims{
		Issuer: "fantasyforum",
		Subject: fmt.Sprintf("%d", user.ID),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, myClaims)
	signedToken, err := token.SignedString([]byte(su.env.TokenSecret))
	return signedToken, err
}
