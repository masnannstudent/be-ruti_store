package service

import (
	"debtomate/module/entities"
	"debtomate/module/feature/auth/domain"
	"debtomate/module/feature/auth/mocks"
	utils "debtomate/utils/mocks"
	"errors"
	"github.com/stretchr/testify/assert"

	"testing"
)

func setupTest(t *testing.T) (*mocks.AuthRepositoryInterface, domain.AuthServiceInterface, *utils.HashInterface, *utils.JWTInterface) {
	repo := mocks.NewAuthRepositoryInterface(t)
	hash := utils.NewHashInterface(t)
	jwt := utils.NewJWTInterface(t)
	service := NewAuthService(repo, hash, jwt)
	return repo, service, hash, jwt
}

func TestLogin(t *testing.T) {
	email := "test@example.com"
	password := "password123"

	t.Run("Success Case - Valid Credentials", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.AdminModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
		}
		expectedToken := "mockedAccessToken"

		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)
		hash.On("ComparePassword", expectedUser.Password, password).Return(true, nil)
		jwt.On("GenerateJWT", expectedUser.ID, expectedUser.Email).Return(expectedToken, nil)

		user, accessToken, err := service.Login(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
		assert.Equal(t, expectedToken, accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})

	t.Run("Error Case - User Not Found", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedErr := errors.New("user not found")
		repo.On("GetUsersByEmail", email).Return(nil, expectedErr)

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertNotCalled(t, "ComparePassword")
		jwt.AssertNotCalled(t, "GenerateJWT")
	})

	t.Run("Error Case - Invalid Credentials", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.AdminModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
		}
		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)

		hash.On("ComparePassword", expectedUser.Password, password).Return(false, nil)

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "wrong credential")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertNotCalled(t, "GenerateJWT")
	})

	t.Run("Error Case - JWT Generation Failure", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.AdminModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
		}
		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)

		hash.On("ComparePassword", expectedUser.Password, password).Return(true, nil)

		jwt.On("GenerateJWT", expectedUser.ID, expectedUser.Email).Return("", errors.New("jwt generation failed"))

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "jwt generation failed")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})
}
