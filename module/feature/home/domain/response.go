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

func ResponseArrayCarousel(data []*entities.CarouselModels) []*CarouselResponse {
	res := make([]*CarouselResponse, 0)

	for _, carouselItem := range data {
		carouselRes := &CarouselResponse{
			ID:        carouselItem.ID,
			Name:      carouselItem.Name,
			Photo:     carouselItem.Photo,
			CreatedAt: carouselItem.CreatedAt,
		}
		res = append(res, carouselRes)
	}

	return res
}
