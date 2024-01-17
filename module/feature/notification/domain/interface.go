package domain

import (
	"ruti-store/module/entities"
)

type NotificationRepositoryInterface interface {
	CreateNotification(notification *entities.NotificationModels) (*entities.NotificationModels, error)
}

type NotificationServiceInterface interface {
	CreateNotification(req *CreateNotificationRequest) (*entities.NotificationModels, error)
}

type NotificationHandlerInterface interface {
}
