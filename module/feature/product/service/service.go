package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/product/domain"
	"time"
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

func (s *ProductService) GetProductByID(productID uint64) (*entities.ProductModels, error) {
	result, err := s.repo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) CreateProduct(req *domain.CreateProductRequest) (*entities.ProductModels, error) {
	newProduct := &entities.ProductModels{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Discount:    req.Discount,
		Stock:       req.Stock,
		CreatedAt:   time.Now(),
	}

	result, err := s.repo.CreateProduct(newProduct, req.CategoryID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProductService) UpdateProduct(productID uint64, req *domain.UpdateProductRequest) error {
	product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	newData := &entities.ProductModels{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Discount:    req.Discount,
		Stock:       req.Stock,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.UpdateProduct(product.ID, newData, req.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
