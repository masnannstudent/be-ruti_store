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

func (r *ReviewRepository) GetReviewsById(reviewID uint64) (*entities.ReviewModels, error) {
	var reviews *entities.ReviewModels

	if err := r.db.Preload("Photos").Where("id = ? AND deleted_at IS NULL", reviewID).
		Order("created_at DESC").
		First(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetPaginatedReviewsByProductID(productID uint64, page, pageSize int) ([]*entities.ReviewModels, error) {
	var reviews []*entities.ReviewModels

	offset := (page - 1) * pageSize

	if err := r.db.
		Preload("User").
		Preload("Photos").
		Where("product_id = ? AND deleted_at IS NULL", productID).
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetTotalReviewsByProductID(productID uint64) (int64, error) {
	var totalReviews int64

	if err := r.db.
		Where("product_id = ? AND deleted_at IS NULL", productID).
		Model(&entities.ReviewModels{}).
		Count(&totalReviews).Error; err != nil {
		return 0, err
	}

	return totalReviews, nil
}

func (r *ReviewRepository) CreateReview(newData *entities.ReviewModels) (*entities.ReviewModels, error) {
	err := r.db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *ReviewRepository) CreateReviewImages(newData *entities.ReviewPhotoModels) (*entities.ReviewPhotoModels, error) {
	err := r.db.Create(newData).Error
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *ReviewRepository) CountAverageRating(productID uint64) (float64, error) {
	var averageRating float64

	query := "SELECT ROUND(AVG(rating), 1) FROM reviews WHERE product_id = ?"
	if err := r.db.Raw(query, productID).Scan(&averageRating).Error; err != nil {
		return 0, err
	}

	return averageRating, nil
}

func (r *ReviewRepository) SetIsReviewed(orderDetailsID, productID uint64) error {
	if err := r.db.Model(&entities.OrderDetailsModels{}).
		Where("id = ? AND product_id = ?", orderDetailsID, productID).
		Update("is_reviewed", true).
		Error; err != nil {
		return err
	}

	return nil
}
