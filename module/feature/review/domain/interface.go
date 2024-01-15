package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ReviewRepositoryInterface interface {
	GetPaginatedReviews(page, pageSize int) ([]*entities.ReviewModels, error)
	GetTotalItems() (int64, error)
	GetReviewsById(reviewID uint64) (*entities.ReviewModels, error)
}

type ReviewServiceInterface interface {
	GetAllReviews(page, pageSize int) ([]*entities.ReviewModels, int64, error)
	GetReviewsPage(currentPage, pageSize int) (int, int, int, int, error)
	GetReviewById(reviewID uint64) (*entities.ReviewModels, error)
}

type ReviewHandlerInterface interface {
	GetAllReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
}
