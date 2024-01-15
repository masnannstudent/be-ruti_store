package domain

import (
	"ruti-store/module/entities"
	"time"
)

type CategoriesResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name" `
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	CreatedAt   time.Time `json:"created_at"`
}

func ResponseArrayCategories(data []*entities.CategoryModels) []*CategoriesResponse {
	res := make([]*CategoriesResponse, 0)

	for _, category := range data {
		categoryRes := &CategoriesResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Photo:       category.Photo,
			CreatedAt:   category.CreatedAt,
		}
		res = append(res, categoryRes)
	}

	return res
}
