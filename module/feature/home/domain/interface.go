package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type HomeRepositoryInterface interface {
	CreateCarousel(carousel *entities.CarouselModels) (*entities.CarouselModels, error)
	GetCarouselById(carouselID uint64) (*entities.CarouselModels, error)
	UpdateCarousel(carouselID uint64, updatedCarousel *entities.CarouselModels) error
	DeleteCarousel(carouselID uint64) error
	GetTotalCarouselItems() (int64, error)
	GetPaginatedCarousel(page, pageSize int) ([]*entities.CarouselModels, error)
	GetTotalProduct() (int64, error)
	GetTotalUser() (int64, error)
	GetTotalIncome() (uint64, error)
	GetAllOrders(page, pageSize int) ([]*entities.OrderModels, error)
	GetTotalOrderItems() (int64, error)
}

type HomeServiceInterface interface {
	CreateCarousel(req *CreateCarouselRequest) (*entities.CarouselModels, error)
	GetCarouselById(carouselID uint64) (*entities.CarouselModels, error)
	UpdateCarousel(carouselID uint64, req *UpdateCarouselRequest) error
	DeleteCarousel(carouselID uint64) error
	GetCarouselPage(currentPage, pageSize int) (int, int, int, int, error)
	GetAllCarouselItems(page, pageSize int) ([]*entities.CarouselModels, int64, error)
	GetDashboardPage() (uint64, int64, int64, error)
	GetAllOrders(page, pageSize int) ([]*entities.OrderModels, int64, error)
	GetOrdersPage(currentPage, pageSize int) (int, int, int, int, error)
}

type HomeHandlerInterface interface {
	CreateCarousel(c *fiber.Ctx) error
	GetCarouselByID(c *fiber.Ctx) error
	GetAllCarouselItems(c *fiber.Ctx) error
	UpdateCarousel(c *fiber.Ctx) error
	DeleteCarousel(c *fiber.Ctx) error
	GetDashboard(c *fiber.Ctx) error
	GetAllOrders(c *fiber.Ctx) error
}
