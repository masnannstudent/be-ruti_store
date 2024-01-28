package domain

import (
	"ruti-store/module/entities"
	"time"
)

type ProductsResponse struct {
	ID           uint64                 `json:"id"`
	Name         string                 `json:"name"`
	Price        uint64                 `json:"price"`
	Description  string                 `json:"description"`
	Discount     uint64                 `json:"discount"`
	Rating       float64                `json:"rating"`
	TotalReviews uint64                 `json:"total_reviews"`
	Size         string                 `json:"size"`
	Color        string                 `json:"color"`
	Stock        uint64                 `json:"stock"`
	Weight       uint64                 `json:"weight"`
	CreatedAt    time.Time              `json:"created_at"`
	Photos       []ProductPhotoResponse `json:"photos"`
}

type ProductPhotoResponse struct {
	ID  uint64 `json:"id"`
	URL string `json:"url"`
}

func ResponseDetailProducts(data *entities.ProductModels) *ProductsResponse {
	res := &ProductsResponse{
		ID:           data.ID,
		Name:         data.Name,
		Price:        data.Price,
		Description:  data.Description,
		Discount:     data.Discount,
		Rating:       data.Rating,
		TotalReviews: data.TotalReviews,
		Size:         data.Size,
		Color:        data.Color,
		Stock:        data.Stock,
		Weight:       data.Weight,
		CreatedAt:    data.CreatedAt,
		Photos:       getPhotoResponses(data.Photos),
	}
	return res
}

func ResponseArrayProducts(data []*entities.ProductModels) []*ProductsResponse {
	res := make([]*ProductsResponse, 0)

	for _, product := range data {
		productRes := &ProductsResponse{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			Description:  product.Description,
			Discount:     product.Discount,
			Rating:       product.Rating,
			TotalReviews: product.TotalReviews,
			Size:         product.Size,
			Color:        product.Color,
			Stock:        product.Stock,
			Weight:       product.Weight,
			CreatedAt:    product.CreatedAt,
			Photos:       getPhotoResponses(product.Photos),
		}
		res = append(res, productRes)
	}

	return res
}

func getPhotoResponses(photos []entities.ProductPhotoModels) []ProductPhotoResponse {
	responses := make([]ProductPhotoResponse, len(photos))
	for i, photo := range photos {
		responses[i] = ProductPhotoResponse{
			ID:  photo.ID,
			URL: photo.URL,
		}
	}
	return responses
}

type ReviewProductFormatter struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Rating      float64 `json:"rating"`
	TotalReview uint64  `json:"total_review"`
}

func ResponseArrayProductReviews(products []*entities.ProductModels) []*ReviewProductFormatter {
	productFormatters := make([]*ReviewProductFormatter, 0)
	for _, product := range products {
		productFormatter := &ReviewProductFormatter{
			ID:          product.ID,
			Name:        product.Name,
			Rating:      product.Rating,
			TotalReview: product.TotalReviews,
		}
		productFormatters = append(productFormatters, productFormatter)
	}
	return productFormatters
}

type AddPhotoProductResponse struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Photo     string `json:"photo"`
}

func ResponseAddPhotoProduct(data *entities.ProductPhotoModels) *AddPhotoProductResponse {
	res := &AddPhotoProductResponse{
		ID:        data.ID,
		ProductID: data.ProductID,
		Photo:     data.URL,
	}
	return res
}
