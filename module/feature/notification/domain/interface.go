package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type NotificationRepositoryInterface interface {
	CreateNotification(notification *entities.NotificationModels) (*entities.NotificationModels, error)
	GetNotificationUser(userID uint64) ([]*entities.NotificationModels, error)
}

type NotificationServiceInterface interface {
	CreateNotification(req *CreateNotificationRequest) (*entities.NotificationModels, error)
	GetNotificationUser(userID uint64) ([]*entities.NotificationModels, error)
}

type NotificationHandlerInterface interface {
	GetNotification(c *fiber.Ctx) error
}
