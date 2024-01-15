package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type ArticleRepositoryInterface interface {
	GetPaginatedArticles(page, pageSize int) ([]*entities.ArticleModels, error)
	GetTotalItems() (int64, error)
	GetArticleByID(articleID uint64) (*entities.ArticleModels, error)
}

type ArticleServiceInterface interface {
	GetAllArticles(page, pageSize int) ([]*entities.ArticleModels, int64, error)
	GetArticlesPage(currentPage, pageSize int) (int, int, int, int, error)
	GetArticleByID(articleID uint64) (*entities.ArticleModels, error)
}

type ArticleHandlerInterface interface {
	GetAllArticles(c *fiber.Ctx) error
	GetArticleByID(c *fiber.Ctx) error
}
