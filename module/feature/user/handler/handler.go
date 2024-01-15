package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
	"ruti-store/module/feature/user/domain"
	"ruti-store/utils/response"
	"strconv"
)

type UserHandler struct {
	service domain.UserServiceInterface
}

func NewUserHandler(service domain.UserServiceInterface) domain.UserHandlerInterface {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetUserByID(userID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved user by ID", domain.UserFormatter(result))
}

func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}
	fmt.Println("User ID: %v", currentUser.ID)
	result, err := h.service.GetUserByID(currentUser.ID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve user profile: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved user profile", domain.UserFormatter(result))
}
