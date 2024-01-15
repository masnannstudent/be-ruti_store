package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/review/domain"
)

type ReviewService struct {
	repo domain.ReviewRepositoryInterface
}

func NewReviewService(repo domain.ReviewRepositoryInterface) domain.ReviewServiceInterface {
	return &ReviewService{
		repo: repo,
	}
}

func (s *ReviewService) GetAllReviews(page, pageSize int) ([]*entities.ReviewModels, int64, error) {
	result, err := s.repo.GetPaginatedReviews(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ReviewService) GetReviewsPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalItems()
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

func (s *ReviewService) GetReviewById(reviewID uint64) (*entities.ReviewModels, error) {
	result, err := s.repo.GetReviewsById(reviewID)
	if err != nil {
		return nil, errors.New("reviews not found")
	}
	return result, nil
}
