package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/home/domain"
	"time"
)

type HomeService struct {
	repo domain.HomeRepositoryInterface
}

func NewHomeService(repo domain.HomeRepositoryInterface) domain.HomeServiceInterface {
	return &HomeService{
		repo: repo,
	}
}

func (s *HomeService) CreateCarousel(req *domain.CreateCarouselRequest) (*entities.CarouselModels, error) {
	newData := &entities.CarouselModels{
		Name:      req.Name,
		Photo:     req.Photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdCategory, err := s.repo.CreateCarousel(newData)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *HomeService) GetCarouselById(carouselID uint64) (*entities.CarouselModels, error) {
	carousels, err := s.repo.GetCarouselById(carouselID)
	if err != nil {
		return nil, errors.New("carousels not found")
	}
	return carousels, nil
}

func (s *HomeService) UpdateCarousel(carouselID uint64, req *domain.UpdateCarouselRequest) error {
	carousels, err := s.repo.GetCarouselById(carouselID)
	if err != nil {
		return errors.New("carousels not found")
	}
	newData := &entities.CarouselModels{
		ID:        carousels.ID,
		Name:      req.Name,
		Photo:     req.Photo,
		UpdatedAt: time.Now(),
	}

	err = s.repo.UpdateCarousel(carousels.ID, newData)
	if err != nil {
		return err
	}
	return nil
}

func (s *HomeService) DeleteCarousel(carouselID uint64) error {
	carousels, err := s.repo.GetCarouselById(carouselID)
	if err != nil {
		return errors.New("carousels not found")
	}
	err = s.repo.DeleteCarousel(carousels.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *HomeService) GetAllCarouselItems(page, pageSize int) ([]*entities.CarouselModels, int64, error) {
	result, err := s.repo.GetPaginatedCarousel(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalCarouselItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *HomeService) GetCarouselPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalCarouselItems()
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

func (s *HomeService) GetDashboardPage() (uint64, int64, int64, error) {
	totalProduct, err := s.repo.GetTotalProduct()
	if err != nil {
		return 0, 0, 0, err
	}
	totalUser, err := s.repo.GetTotalUser()
	if err != nil {
		return 0, 0, 0, err
	}

	totalIncome, err := s.repo.GetTotalIncome()
	if err != nil {
		return 0, 0, 0, err
	}

	return totalIncome, totalProduct, totalUser, nil

}
