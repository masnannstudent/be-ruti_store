package domain

import (
	"ruti-store/module/entities"
	"time"
)

type ReviewPhotoResponse struct {
	ID        uint64    `json:"id"`
	ReviewID  uint64    `json:"review_id"`
	ImageURL  string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewResponse struct {
	ID          uint64                `json:"id"`
	UserID      uint64                `json:"user_id"`
	ProductID   uint64                `json:"product_id"`
	Rating      uint64                `json:"rating"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `gorm:"column:created_at;type:timestamp" json:"created_at"`
	Photos      []ReviewPhotoResponse `gorm:"foreignKey:ReviewID" json:"photos"`
}

func ResponseArrayReviews(data []*entities.ReviewModels) []*ReviewResponse {
	res := make([]*ReviewResponse, 0)

	for _, review := range data {
		reviewRes := &ReviewResponse{
			ID:          review.ID,
			UserID:      review.UserID,
			ProductID:   review.ProductID,
			Rating:      review.Rating,
			Description: review.Description,
			CreatedAt:   review.CreatedAt,
		}
		res = append(res, reviewRes)
	}

	return res
}

func ReviewFormatter(review *entities.ReviewModels) *ReviewResponse {
	return &ReviewResponse{
		ID:          review.ID,
		UserID:      review.UserID,
		ProductID:   review.ProductID,
		Rating:      review.Rating,
		Description: review.Description,
		CreatedAt:   review.CreatedAt,
	}
}
