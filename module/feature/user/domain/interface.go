package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type UserRepositoryInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
	EditProfile(userID uint64, req *entities.UserModels) error
}

type UserServiceInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
	EditProfile(userID uint64, req *EditProfileRequest) error
}

type UserHandlerInterface interface {
	GetUserByID(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
	EditProfile(c *fiber.Ctx) error
}
