package domain

import (
	"ruti-store/module/entities"
	"time"
)

type CarouselResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
}

func CarouselFormatter(carousel *entities.CarouselModels) *CarouselResponse {
	carouselFormatter := &CarouselResponse{
		ID:        carousel.ID,
		Name:      carousel.Name,
		Photo:     carousel.Photo,
		CreatedAt: carousel.CreatedAt,
	}
	return carouselFormatter
}
