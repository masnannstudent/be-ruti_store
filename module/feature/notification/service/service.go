package service

import (
	"ruti-store/module/entities"
	"ruti-store/module/feature/notification/domain"
	"time"
)

type NotificationService struct {
	repo domain.NotificationRepositoryInterface
}

func NewNotificationService(repo domain.NotificationRepositoryInterface) domain.NotificationServiceInterface {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) CreateNotification(req *domain.CreateNotificationRequest) (*entities.NotificationModels, error) {
	newData := &entities.NotificationModels{
		UserID:    req.UserID,
		OrderID:   req.OrderID,
		Title:     req.Title,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}
	result, err := s.repo.CreateNotification(newData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NotificationService) GetNotificationUser(userID uint64) ([]*entities.NotificationModels, error) {
	result, err := s.repo.GetNotificationUser(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
