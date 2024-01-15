package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/review/domain"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) domain.ReviewRepositoryInterface {
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ReviewModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ReviewRepository) GetPaginatedReviews(page, pageSize int) ([]*entities.ReviewModels, error) {
	var reviews []*entities.ReviewModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReviewsById(reviewID uint64) (*entities.ReviewModels, error) {
	var reviews *entities.ReviewModels

	if err := r.db.Preload("Photos").Where("id = ? AND deleted_at IS NULL", reviewID).First(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}
