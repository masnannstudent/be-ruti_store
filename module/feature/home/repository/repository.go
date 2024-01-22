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

func (r *HomeRepository) GetTotalCarouselItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.CarouselModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *HomeRepository) GetPaginatedCarousel(page, pageSize int) ([]*entities.CarouselModels, error) {
	var carousels []*entities.CarouselModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Find(&carousels).
		Order("created_at DESC").
		Error; err != nil {
		return nil, err
	}

	return carousels, nil
}

func (r *HomeRepository) GetTotalProduct() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ProductModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *HomeRepository) GetTotalUser() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.UserModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *HomeRepository) GetTotalIncome() (uint64, error) {
	var totalIncome uint64

	err := r.db.Model(&entities.OrderModels{}).
		Select("SUM(total_amount_paid) as total_income").
		Where("payment_status = ?", "Konfirmasi").
		Pluck("total_income", &totalIncome).
		Error

	if err != nil {
		return 0, err
	}

	return totalIncome, nil
}

func (r *HomeRepository) GetAllOrders(page, pageSize int) ([]*entities.OrderModels, error) {
	var orders []*entities.OrderModels

	offset := (page - 1) * pageSize

	if err := r.db.
		Preload("User").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *HomeRepository) GetTotalOrderItems() (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.OrderModels{}).Count(&totalItems).Where("deleted_at IS NULL").Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}
