package service

import (
	"errors"
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

func (s *HomeService) UpdateCarousel(carouselID uint64, updatedCarousel *entities.CarouselModels) error {
	carousels, err := s.repo.GetCarouselById(carouselID)
	if err != nil {
		return errors.New("carousels not found")
	}
	err = s.repo.UpdateCarousel(carousels.ID, updatedCarousel)
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
