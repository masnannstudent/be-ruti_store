package domain

import (
	"ruti-store/module/entities"
	"time"
)

type ArticleResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
}

func ResponseArrayArticles(data []*entities.ArticleModels) []*ArticleResponse {
	res := make([]*ArticleResponse, 0)

	for _, article := range data {
		articleRes := &ArticleResponse{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Author:    article.Author,
			Photo:     article.Photo,
			CreatedAt: article.CreatedAt,
		}
		res = append(res, articleRes)
	}

	return res
}

func ArticleDetailFormatter(article *entities.ArticleModels) *ArticleResponse {
	articles := &ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Author:    article.Author,
		Photo:     article.Photo,
		CreatedAt: article.CreatedAt,
	}
	return articles
}
