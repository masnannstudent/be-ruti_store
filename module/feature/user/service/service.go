package service

import (
	"errors"
	"math"
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

func (s *UserService) GetAllUserItems(page, pageSize int) ([]*entities.UserModels, int64, error) {
	result, err := s.repo.GetPaginatedUsers(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalUserItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *UserService) GetUserPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalUserItems()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}
