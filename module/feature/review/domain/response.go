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
	User        UserResponse          `json:"user"`
	ProductID   uint64                `json:"product_id"`
	Rating      uint64                `json:"rating"`
	Description string                `json:"description"`
	CreatedAt   time.Time             `json:"created_at"`
	Photos      []ReviewPhotoResponse `json:"photos"`
}

type UserResponse struct {
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
}

func ResponseArrayReviews(data []*entities.ReviewModels) []*ReviewResponse {
	res := make([]*ReviewResponse, 0)

	for _, review := range data {
		reviewRes := &ReviewResponse{
			ID:     review.ID,
			UserID: review.UserID,
			User: UserResponse{
				Name:         review.User.Name,
				PhotoProfile: review.User.PhotoProfile,
			},
			ProductID:   review.ProductID,
			Rating:      review.Rating,
			Description: review.Description,
			CreatedAt:   review.CreatedAt,
			Photos:      FormatReviewPhotos(review.Photos),
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
		Photos:      FormatReviewPhotos(review.Photos),
	}
}

func FormatReviewPhotos(photos []entities.ReviewPhotoModels) []ReviewPhotoResponse {
	formattedPhotos := make([]ReviewPhotoResponse, len(photos))
	for i, photo := range photos {
		formattedPhotos[i] = ReviewPhotoResponse{
			ID:        photo.ID,
			ReviewID:  photo.ReviewID,
			ImageURL:  photo.ImageURL,
			CreatedAt: photo.CreatedAt,
		}
	}
	return formattedPhotos
}

type ReviewPhotosFormatter struct {
	ID       uint64 `json:"id"`
	ImageURL string `json:"url"`
}

func FormatCreateReviewPhotos(photo *entities.ReviewPhotoModels) *ReviewPhotosFormatter {
	formattedPhoto := &ReviewPhotosFormatter{
		ID:       photo.ID,
		ImageURL: photo.ImageURL,
	}

	return formattedPhoto
}
