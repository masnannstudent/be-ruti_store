package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type UserRepositoryInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
	EditProfile(userID uint64, req *entities.UserModels) error
	GetTotalUserItems() (int64, error)
	GetPaginatedUsers(page, pageSize int) ([]*entities.UserModels, error)
	ChatBotAI(req *CreateChatBotRequest) (string, error)
}

type UserServiceInterface interface {
	GetUserByID(userID uint64) (*entities.UserModels, error)
	EditProfile(userID uint64, req *EditProfileRequest) error
	GetAllUserItems(page, pageSize int) ([]*entities.UserModels, int64, error)
	GetUserPage(currentPage, pageSize int) (int, int, int, int, error)
	ChatBot(req *CreateChatBotRequest) (string, error)
}

type UserHandlerInterface interface {
	GetUserByID(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
	EditProfile(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
	ChatBot(c *fiber.Ctx) error
}
