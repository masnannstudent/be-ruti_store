package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/category/domain"
	"time"
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

func (s *CategoryService) GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error) {
	result, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return result, nil
}

func (s *CategoryService) CreateCategory(req *domain.CreateCategoryRequest) (*entities.CategoryModels, error) {
	newData := &entities.CategoryModels{
		Name:        req.Name,
		Description: req.Description,
		Photo:       req.Photo,
		CreatedAt:   time.Now(),
	}

	createdCategory, err := s.repo.CreateCategory(newData)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *CategoryService) UpdateCategory(categoryID uint64, req *domain.UpdateCategoryRequest) error {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	newData := &entities.CategoryModels{
		ID:          category.ID,
		Name:        req.Name,
		Description: req.Description,
		Photo:       req.Photo,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.UpdateCategory(category.ID, newData)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteCategory(categoryID uint64) error {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return errors.New("category not found")
	}

	err = s.repo.DeleteCategory(category.ID)
	if err != nil {
		return err
	}

	return nil
}
