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
	CreateCart(newCart *entities.CartModels) (*entities.CartModels, error)
	GetCartItem(userID, productID uint64) (*entities.CartModels, error)
	UpdateCartItem(cartItem *entities.CartModels) error
	GetCartByID(cartID uint64) (*entities.CartModels, error)
	DeleteCartItem(cartItemID uint64) error
	GetCartByUserID(userID uint64) ([]*entities.CartModels, error)
	AcceptOrder(orderID, orderStatus string) error
}

type OrderServiceInterface interface {
	GetAllOrders(page, pageSize int) ([]*entities.OrderModels, int64, error)
	GetOrdersPage(currentPage, pageSize int) (int, int, int, int, error)
	CreateOrder(userID uint64, request *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrderByID(orderID string) (*entities.OrderModels, error)
	CallBack(req map[string]interface{}) error
	CreateCart(userID uint64, req *CreateCartRequest) (*entities.CartModels, error)
	DeleteCartItems(cartID uint64) error
	GetCartUser(userID uint64) ([]*entities.CartModels, error)
	CreateOrderCart(userID uint64, request *CreateOrderCartRequest) (*CreateOrderResponse, error)
	AcceptOrder(orderID string) error
}

type OrderHandlerInterface interface {
	GetAllOrders(c *fiber.Ctx) error
	GetAllPayment(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	Callback(c *fiber.Ctx) error
	CreateCart(c *fiber.Ctx) error
	DeleteCart(c *fiber.Ctx) error
	GetCartUser(c *fiber.Ctx) error
	CreateOrderCart(c *fiber.Ctx) error
	AcceptOrder(c *fiber.Ctx) error
}
