package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type OrderRepositoryInterface interface {
	GetTotalItems() (int64, error)
	GetPaginatedOrders(page, pageSize int) ([]*entities.OrderModels, error)
}

type OrderServiceInterface interface {
	GetAllOrders(page, pageSize int) ([]*entities.OrderModels, int64, error)
	GetOrdersPage(currentPage, pageSize int) (int, int, int, int, error)
}

type OrderHandlerInterface interface {
	GetAllOrders(c *fiber.Ctx) error
	GetAllPayment(c *fiber.Ctx) error
}
