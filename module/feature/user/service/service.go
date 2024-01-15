package service

import (
	"errors"
	"ruti-store/module/entities"
	"ruti-store/module/feature/user/domain"
	"time"
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

func (s *UserService) EditProfile(userID uint64, req *domain.EditProfileRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	newData := &entities.UserModels{
		Phone:        req.Phone,
		Name:         req.Name,
		PhotoProfile: req.PhotoProfile,
		UpdatedAt:    time.Now(),
	}
	err = s.repo.EditProfile(user.ID, newData)
	if err != nil {
		return err
	}

	return nil
}
