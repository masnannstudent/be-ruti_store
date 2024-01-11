package service

import (
	"errors"
	"ruti-store/module/entities"
	"ruti-store/module/feature/auth/domain"
	"ruti-store/utils/hash"
	"ruti-store/utils/token"
)

type AuthService struct {
	repo domain.AuthRepositoryInterface
	hash hash.HashInterface
	jwt  token.JWTInterface
}

func NewAuthService(
	repo domain.AuthRepositoryInterface,
	hash hash.HashInterface,
	jwt token.JWTInterface,
) domain.AuthServiceInterface {
	return &AuthService{
		repo: repo,
		hash: hash,
		jwt:  jwt,
	}
}

func (s *AuthService) Login(email, password string) (*entities.UserModels, string, error) {

	user, err := s.repo.GetUsersByEmail(email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	isValidPassword, err := s.hash.ComparePassword(user.Password, password)
	if err != nil || !isValidPassword {
		return nil, "", errors.New("wrong credential")
	}

	accessToken, err := s.jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, accessToken, nil
}
