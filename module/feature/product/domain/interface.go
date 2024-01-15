package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ProductRepositoryInterface interface {
	GetPaginatedProducts(page, pageSize int) ([]*entities.ProductModels, error)
	GetTotalItems() (int64, error)
	GetProductByID(productID uint64) (*entities.ProductModels, error)
	CreateProduct(product *entities.ProductModels, categoryIDs []uint64) (*entities.ProductModels, error)
	UpdateProduct(productID uint64, newData *entities.ProductModels, categoryIDs []uint64) error
}

type ProductServiceInterface interface {
	GetAllProducts(page, pageSize int) ([]*entities.ProductModels, int64, error)
	GetProductsPage(currentPage, pageSize int) (int, int, int, int, error)
	GetProductByID(productID uint64) (*entities.ProductModels, error)
	CreateProduct(req *CreateProductRequest) (*entities.ProductModels, error)
	UpdateProduct(productID uint64, req *UpdateProductRequest) error
}

type ProductHandlerInterface interface {
	GetAllProducts(c *fiber.Ctx) error
	GetProductByID(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
}
