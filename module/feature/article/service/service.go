package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/article/domain"
)

type ArticleService struct {
	repo domain.ArticleRepositoryInterface
}

func NewArticleService(repo domain.ArticleRepositoryInterface) domain.ArticleServiceInterface {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetAllArticles(page, pageSize int) ([]*entities.ArticleModels, int64, error) {
	result, err := s.repo.GetPaginatedArticles(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ArticleService) GetArticlesPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}

func (s *ArticleService) GetArticleByID(articleID uint64) (*entities.ArticleModels, error) {
	result, err := s.repo.GetArticleByID(articleID)
	if err != nil {
		return nil, errors.New("article not found")
	}
	return result, nil
}
