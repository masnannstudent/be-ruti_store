package handler

import (
	"debtomate/module/feature/auth/domain"
	"debtomate/utils/response"
	"debtomate/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service domain.AuthServiceInterface
}

func NewAuthHandler(service domain.AuthServiceInterface) domain.AuthHandlerInterface {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(domain.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "err")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Status internal server error: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Login successfully", domain.LoginFormatter(user, token))
}
