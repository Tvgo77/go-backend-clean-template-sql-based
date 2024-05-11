package usecase

import (
	"context"
	"fmt"
	"go-backend/domain"
	"go-backend/setup"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	env *setup.Env
}

func NewLoginUsecase(ur domain.UserRepository, env *setup.Env) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		env: env,
	}
}

func (lu *loginUsecase) HasUser(ctx context.Context, email string) (bool, error) {
	hasUser, err := lu.userRepository.CheckExistByEmail(ctx, email)
	return hasUser, err
}

func (lu *loginUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	conds := &domain.User{Email: email}
	user, err := lu.userRepository.Fetch(ctx, conds)
	return user, err
}

func (lu *loginUsecase) NewJWTtoken(user *domain.User) (string, error) {
	myClaims := jwt.RegisteredClaims{
		Issuer: "fantasyforum",
		Subject: fmt.Sprintf("%d", user.ID),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	signedToken, err := token.SignedString([]byte(lu.env.TokenSecret))
	return signedToken, err
}