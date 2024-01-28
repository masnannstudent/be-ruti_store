package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ReviewRepositoryInterface interface {
	GetReviewsById(reviewID uint64) (*entities.ReviewModels, error)
	GetPaginatedReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, error)
	GetTotalReviewsByProductID(productID uint64) (int64, error)
	CreateReview(newData *entities.ReviewModels) (*entities.ReviewModels, error)
	CreateReviewImages(newData *entities.ReviewPhotoModels) (*entities.ReviewPhotoModels, error)
	CountAverageRating(productID uint64) (float64, error)
	SetIsReviewed(orderDetailsID, productID uint64) error
}

type ReviewServiceInterface interface {
	GetReviewById(reviewID uint64) (*entities.ReviewModels, error)
	GetReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, int64, error)
	GetReviewsProductPage(productID uint64, currentPage, pageSize int) (int, int, int, int, error)
	CreateReview(userID uint64, req *CreateReviewRequest) (*entities.ReviewModels, error)
	CreateReviewImages(req *CreatePhotoReviewRequest) (*entities.ReviewPhotoModels, error)
}

type ReviewHandlerInterface interface {
	GetReviewByID(c *fiber.Ctx) error
	GetAllReviewProduct(c *fiber.Ctx) error
	CreateReview(c *fiber.Ctx) error
	CreateReviewPhoto(c *fiber.Ctx) error
}
