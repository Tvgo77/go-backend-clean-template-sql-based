package usecase

import (
	"context"
	"errors"
	mock "go-backend/domain/mock"
	"go-backend/setup"
	"go-backend/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


// This test is not necessary. Just a simple practice for mock and unit test
func TestHasUser(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		// Step 1: Create mock controller
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// Step 2: Create mock object and normal object
		mur := mock.NewMockUserRepository(mockCtrl)
		env := setup.NewEnv()

		// Step 3: Specify expecting para and return value of mock obj
		mur.EXPECT().CheckExistByEmail(gomock.Any(), gomock.AssignableToTypeOf(string(""))).Return(true, nil)

		// Step 4: Creat obj to be test
		su := NewSignupUsecase(mur, env)
		hasUser, err := su.HasUser(context.Background(), "test@gmail.com")

		assert.True(t, hasUser)
		assert.NoError(t, err)
	})

	t.Run("Not Found", func(t *testing.T) {
		// Step 1: Create mock controller
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// Step 2: Create mock object and normal object
		mur := mock.NewMockUserRepository(mockCtrl)
		env := setup.NewEnv()

		// Step 3: Specify expecting para and return value of mock obj
		mur.EXPECT().CheckExistByEmail(gomock.Any(), gomock.AssignableToTypeOf(string(""))).Return(false, nil)

		// Step 4: Creat obj to be test
		su := NewSignupUsecase(mur, env)
		hasUser, err := su.HasUser(context.Background(), "test@gmail.com")

		assert.False(t, hasUser)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		// Step 1: Create mock controller
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// Step 2: Create mock object and normal object
		mur := mock.NewMockUserRepository(mockCtrl)
		env := setup.NewEnv()

		// Step 3: Specify expecting para and return value of mock obj
		mur.EXPECT().CheckExistByEmail(gomock.Any(), gomock.AssignableToTypeOf(string(""))).Return(false, errors.New(""))

		// Step 4: Creat obj to be test
		su := NewSignupUsecase(mur, env)
		hasUser, err := su.HasUser(context.Background(), "test@gmail.com")

		assert.False(t, hasUser)
		assert.Error(t, err)
	})
}

func TestNewJWTtoken(t *testing.T) {
	su := signupUsecase{env: setup.NewEnv()}
	_, err := su.NewJWTtoken(&domain.User{ID: 1})
	assert.NoError(t, err)
}
