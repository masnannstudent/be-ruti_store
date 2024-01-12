package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ProductRepositoryInterface interface {
	GetPaginatedProducts(page, pageSize int) ([]*entities.ProductModels, error)
	GetTotalItems() (int64, error)
}

type ProductServiceInterface interface {
	GetAllProducts(page, pageSize int) ([]*entities.ProductModels, int64, error)
	GetProductsPage(currentPage, pageSize int) (int, int, int, int, error)
}

type ProductHandlerInterface interface {
	GetAllProducts(c *fiber.Ctx) error
}
