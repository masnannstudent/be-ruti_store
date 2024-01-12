package service

import (
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/product/domain"
)

type ProductService struct {
	repo domain.ProductRepositoryInterface
}

func NewProductService(repo domain.ProductRepositoryInterface) domain.ProductServiceInterface {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAllProducts(page, pageSize int) ([]*entities.ProductModels, int64, error) {
	result, err := s.repo.GetPaginatedProducts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ProductService) GetProductsPage(currentPage, pageSize int) (int, int, int, int, error) {
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
