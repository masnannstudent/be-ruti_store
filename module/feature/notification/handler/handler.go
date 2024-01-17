package handler

import "ruti-store/module/feature/notification/domain"

type NotificationHandler struct {
	service domain.NotificationServiceInterface
}

func NewNotificationHandler(service domain.NotificationServiceInterface) domain.NotificationHandlerInterface {
	return &NotificationHandler{
		service: service,
	}
}
