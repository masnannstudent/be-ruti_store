package service

import (
	"ruti-store/module/entities"
	"ruti-store/module/feature/user/domain"
)

type UserService struct {
	repo domain.UserRepositoryInterface
}

func NewUserService(repo domain.UserRepositoryInterface) domain.UserServiceInterface {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(userID uint64) (*entities.UserModels, error) {
	result, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
