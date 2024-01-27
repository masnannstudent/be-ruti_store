package domain

type CreateReviewRequest struct {
	ProductID      uint64 `json:"product_id" validate:"required"`
	OrderDetailsID uint64 `json:"order_details_id" validate:"required"`
	Rating         uint64 `json:"rating" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type CreatePhotoReviewRequest struct {
	ReviewID uint64 `form:"review_id" json:"review_id"`
	Photo    string `form:"photo" json:"photo"`
}
