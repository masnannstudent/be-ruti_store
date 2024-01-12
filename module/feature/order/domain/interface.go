package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go/snap"
	"ruti-store/module/entities"
)

type OrderRepositoryInterface interface {
	GetTotalItems() (int64, error)
	GetPaginatedOrders(page, pageSize int) ([]*entities.OrderModels, error)
	CreateOrder(newOrder *entities.OrderModels) (*entities.OrderModels, error)
	CreateSnap(orderID, name, email string, totalAmountPaid uint64) (*snap.Response, error)
	CheckTransaction(orderID string) (Status, error)
	GetOrderByID(orderID string) (*entities.OrderModels, error)
	UpdatePayment(orderID, orderStatus, paymentStatus string) error
}

type OrderServiceInterface interface {
	GetAllOrders(page, pageSize int) ([]*entities.OrderModels, int64, error)
	GetOrdersPage(currentPage, pageSize int) (int, int, int, int, error)
	CreateOrder(userID uint64, request *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrderByID(orderID string) (*entities.OrderModels, error)
	CallBack(req map[string]interface{}) error
}

type OrderHandlerInterface interface {
	GetAllOrders(c *fiber.Ctx) error
	GetAllPayment(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	Callback(c *fiber.Ctx) error
}
