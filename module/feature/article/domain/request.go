package domain

type CreateArticleRequest struct {
	Title   string `form:"title" validate:"required"`
	Content string `form:"content" validate:"required"`
	Photo   string `form:"photo"`
}

type UpdateArticleRequest struct {
	Title   string `form:"title" `
	Content string `form:"content"`
	Photo   string `form:"photo"`
}
