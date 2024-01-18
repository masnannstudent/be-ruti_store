package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/notification/domain"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepositoryInterface {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) CreateNotification(notification *entities.NotificationModels) (*entities.NotificationModels, error) {
	err := r.db.Create(notification).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) GetNotificationUser(userID uint64) ([]*entities.NotificationModels, error) {
	var notify []*entities.NotificationModels
	if err := r.db.Where("user_id", userID).
		Order("created_at DESC").
		Find(&notify).Error; err != nil {
		return nil, err
	}
	return notify, nil
}
