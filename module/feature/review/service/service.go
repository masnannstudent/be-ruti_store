package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	product "ruti-store/module/feature/product/domain"
	"ruti-store/module/feature/review/domain"
	"time"
)

type ReviewService struct {
	repo           domain.ReviewRepositoryInterface
	productService product.ProductServiceInterface
}

func NewReviewService(repo domain.ReviewRepositoryInterface, productService product.ProductServiceInterface) domain.ReviewServiceInterface {
	return &ReviewService{
		repo:           repo,
		productService: productService,
	}
}

func (s *ReviewService) GetReviewById(reviewID uint64) (*entities.ReviewModels, error) {
	result, err := s.repo.GetReviewsById(reviewID)
	if err != nil {
		return nil, errors.New("reviews not found")
	}
	return result, nil
}

func (s *ReviewService) GetReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, int64, error) {
	reviews, err := s.repo.GetPaginatedReviewsByProductID(productID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalReviews, err := s.repo.GetTotalReviewsByProductID(productID)
	if err != nil {
		return nil, 0, err
	}

	return reviews, totalReviews, nil
}

func (s *ReviewService) GetReviewsProductPage(productID uint64, currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalReviewsByProductID(productID)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}

func (s *ReviewService) CreateReview(userID uint64, req *domain.CreateReviewRequest) (*entities.ReviewModels, error) {
	products, err := s.productService.GetProductByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if req.Rating > 5 {
		return nil, errors.New("rating should not exceed 5")
	}

	value := &entities.ReviewModels{
		UserID:         userID,
		ProductID:      products.ID,
		OrderDetailsID: req.OrderDetailsID,
		Rating:         req.Rating,
		Description:    req.Description,
		CreatedAt:      time.Now(),
	}

	createdReview, err := s.repo.CreateReview(value)
	if err != nil {
		return nil, err
	}

	err = s.productService.UpdateTotalReview(createdReview.ProductID)
	if err != nil {
		return nil, errors.New("failed to update total reviews")
	}

	averageRating, err := s.repo.CountAverageRating(createdReview.ProductID)
	if err != nil {
		return nil, errors.New("failed to calculate product average rating")
	}

	err = s.productService.UpdateProductRating(createdReview.ProductID, averageRating)
	if err != nil {
		return nil, errors.New("failed to update product rating")
	}

	err = s.repo.SetIsReviewed(req.OrderDetailsID, createdReview.ProductID)
	if err != nil {
		return nil, errors.New("failed to set is_reviewed in order details")
	}

	return createdReview, nil
}

func (s *ReviewService) CreateReviewImages(req *domain.CreatePhotoReviewRequest) (*entities.ReviewPhotoModels, error) {
	review, err := s.repo.GetReviewsById(req.ReviewID)
	if err != nil {
		return nil, errors.New("review not found")
	}
	value := &entities.ReviewPhotoModels{
		ReviewID:  review.ID,
		ImageURL:  req.Photo,
		CreatedAt: time.Now(),
	}

	createdReviewPhoto, err := s.repo.CreateReviewImages(value)
	if err != nil {
		return nil, err
	}

	return createdReviewPhoto, nil
}
