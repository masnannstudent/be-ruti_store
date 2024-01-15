package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type CategoryRepositoryInterface interface {
	GetPaginatedCategories(page, pageSize int) ([]*entities.CategoryModels, error)
	GetTotalItems() (int64, error)
	GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error)
}

type CategoryServiceInterface interface {
	GetAllCategories(page, pageSize int) ([]*entities.CategoryModels, int64, error)
	GetCategoriesPage(currentPage, pageSize int) (int, int, int, int, error)
	GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error)
}

type CategoryHandlerInterface interface {
	GetAllCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
}
