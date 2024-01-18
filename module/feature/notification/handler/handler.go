package handler

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
	"ruti-store/module/feature/notification/domain"
	"ruti-store/utils/response"
)

type NotificationHandler struct {
	service domain.NotificationServiceInterface
}

func NewNotificationHandler(service domain.NotificationServiceInterface) domain.NotificationHandlerInterface {
	return &NotificationHandler{
		service: service,
	}
}

func (h *NotificationHandler) GetNotification(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	result, err := h.service.GetNotificationUser(currentUser.ID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Success get notification user", domain.ResponseArrayNotificationUser(result))

}
