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

type DashboardResponse struct {
	TotalIncome  uint64 `json:"total_income"`
	TotalProduct int64  `json:"total_product"`
	TotalUser    int64  `json:"total_user"`
}

func FormatDashboardResponse(totalIncome uint64, totalProduct, totalUser int64) *DashboardResponse {
	result := &DashboardResponse{
		TotalIncome:  totalIncome,
		TotalProduct: totalProduct,
		TotalUser:    totalUser,
	}
	return result
}
