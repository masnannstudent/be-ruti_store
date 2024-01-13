package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/home/domain"
	"time"
)

type HomeRepository struct {
	db *gorm.DB
}

func NewHomeRepository(db *gorm.DB) domain.HomeRepositoryInterface {
	return &HomeRepository{
		db: db,
	}
}

func (r *HomeRepository) CreateCarousel(carousel *entities.CarouselModels) (*entities.CarouselModels, error) {
	err := r.db.Create(&carousel).Error
	if err != nil {
		return nil, err
	}
	return carousel, nil
}

func (r *HomeRepository) GetCarouselById(carouselID uint64) (*entities.CarouselModels, error) {
	var carousels *entities.CarouselModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", carouselID).First(&carousels).Error; err != nil {
		return nil, err
	}
	return carousels, nil
}

func (r *HomeRepository) UpdateCarousel(carouselID uint64, updatedCarousel *entities.CarouselModels) error {
	var carousels *entities.CarouselModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", carouselID).First(&carousels).Error; err != nil {
		return err
	}
	if err := r.db.Updates(&updatedCarousel).Error; err != nil {
		return err
	}
	return nil
}

func (r *HomeRepository) DeleteCarousel(carouselID uint64) error {
	carousels := &entities.CarouselModels{}
	if err := r.db.First(carousels, carouselID).Error; err != nil {
		return err
	}

	if err := r.db.Model(carousels).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
