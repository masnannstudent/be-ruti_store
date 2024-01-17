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
		Weight:      req.Weight,
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
		Weight:      req.Weight,
		UpdatedAt:   time.Now(),
	}

	err = s.repo.UpdateProduct(product.ID, newData, req.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductService) DeleteProduct(productID uint64) error {
	product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	err = s.repo.DeleteProduct(product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateTotalReview(productID uint64) error {
	products, err := s.repo.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}
	err = s.repo.UpdateTotalReview(products.ID)
	if err != nil {
		return errors.New("failed to update total reviews")
	}

	return nil
}

func (s *ProductService) UpdateProductRating(productID uint64, newRating float64) error {
	err := s.repo.UpdateProductRating(productID, newRating)
	if err != nil {
		return errors.New("failed to update product rating")
	}

	return nil
}

func (s *ProductService) GetProductReviews(page, perPage int) ([]*entities.ProductModels, int64, error) {
	products, err := s.repo.GetProductReviews(page, perPage)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return products, totalItems, nil
}
