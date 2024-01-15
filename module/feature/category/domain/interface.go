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
}

type CategoryServiceInterface interface {
	GetAllCategories(page, pageSize int) ([]*entities.CategoryModels, int64, error)
	GetCategoriesPage(currentPage, pageSize int) (int, int, int, int, error)
	GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error)
	CreateCategory(req *CreateCategoryRequest) (*entities.CategoryModels, error)
	UpdateCategory(categoryID uint64, req *UpdateCategoryRequest) error
}

type CategoryHandlerInterface interface {
	GetAllCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
}
