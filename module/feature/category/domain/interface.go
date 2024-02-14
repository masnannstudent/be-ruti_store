package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type CategoryRepositoryInterface interface {
	GetPaginatedCategories(page, pageSize int) ([]*entities.CategoryModels, error)
	GetTotalItems() (int64, error)
	GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error)
	CreateCategory(category *entities.CategoryModels) (*entities.CategoryModels, error)
	UpdateCategory(categoryID uint64, updatedCategory *entities.CategoryModels) error
	DeleteCategory(categoryID uint64) error
	GetProductsByCategoryID(page, perPage int, categoryID uint64) ([]*entities.ProductModels, int64, error)
}

type CategoryServiceInterface interface {
	GetAllCategories(page, pageSize int) ([]*entities.CategoryModels, int64, error)
	GetCategoryPage(currentPage, pageSize, totalItems int) (int, int, int, error)
	GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error)
	CreateCategory(req *CreateCategoryRequest) (*entities.CategoryModels, error)
	UpdateCategory(categoryID uint64, req *UpdateCategoryRequest) error
	DeleteCategory(categoryID uint64) error
	SearchProductByCategoryID(page, pageSize int, categoryID uint64) ([]*entities.ProductModels, int64, error)
}

type CategoryHandlerInterface interface {
	GetAllCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
	GetAllProductByCategoryID(c *fiber.Ctx) error
}
