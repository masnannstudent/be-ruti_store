package domain

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Photo       string `json:"photo"`
}
