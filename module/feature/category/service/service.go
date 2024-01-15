package service

import (
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/category/domain"
)

type CategoryService struct {
	repo domain.CategoryRepositoryInterface
}

func NewCategoryService(repo domain.CategoryRepositoryInterface) domain.CategoryServiceInterface {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetAllCategories(page, pageSize int) ([]*entities.CategoryModels, int64, error) {
	result, err := s.repo.GetPaginatedCategories(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *CategoryService) GetCategoriesPage(currentPage, pageSize int) (int, int, int, int, error) {
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
