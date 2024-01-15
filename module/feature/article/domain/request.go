package domain

type CreateArticleRequest struct {
	Title   string `form:"title" validate:"required"`
	Content string `form:"content" validate:"required"`
	Photo   string `form:"photo"`
}
