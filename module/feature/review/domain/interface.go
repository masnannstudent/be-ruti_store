package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ReviewRepositoryInterface interface {
	GetPaginatedReviews(page, pageSize int) ([]*entities.ReviewModels, error)
	GetTotalItems() (int64, error)
	GetReviewsById(reviewID uint64) (*entities.ReviewModels, error)
	GetPaginatedReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, error)
	GetTotalReviewsByProductID(productID uint64) (int64, error)
}

type ReviewServiceInterface interface {
	GetAllReviews(page, pageSize int) ([]*entities.ReviewModels, int64, error)
	GetReviewsPage(currentPage, pageSize int) (int, int, int, int, error)
	GetReviewById(reviewID uint64) (*entities.ReviewModels, error)
	GetReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, int64, error)
	GetReviewsProductPage(productID uint64, currentPage, pageSize int) (int, int, int, int, error)
}

type ReviewHandlerInterface interface {
	GetAllReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
	GetAllReviewProduct(c *fiber.Ctx) error
}
