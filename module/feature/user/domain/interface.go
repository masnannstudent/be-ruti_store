package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type UserRepositoryInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserServiceInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
}

type UserHandlerInterface interface {
	GetUserByID(c *fiber.Ctx) error
}
